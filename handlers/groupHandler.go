package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/models"
	"todo/modelsDTO"
	"todo/services"
)

type GroupHandler interface {
	GroupHandle(w http.ResponseWriter, r *http.Request) error
}

type groupHandler struct {
	groupService services.GroupService
}

func NewGroupHandler(groupService services.GroupService) GroupHandler {
	return &groupHandler{groupService: groupService}
}

func (groupHandler *groupHandler) GroupHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		var group models.Group
		err := json.NewDecoder(r.Body).Decode(&group)
		if err != nil {
			return err
		}

		id, err := groupHandler.groupService.AddGroupS(group)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		return json.NewEncoder(w).Encode(id)
	}
	if r.Method == http.MethodGet {
		userIdStr := r.URL.Query().Get("user_id")

		if userIdStr != "" {
			id, err := strconv.Atoi(userIdStr)
			if err != nil {
				return err
			}
			errTwo := json.NewEncoder(w).Encode(groupHandler.groupService.GetGroupsByUserS(id))
			if errTwo != nil {
				return errTwo
			}
		} else {
			err := json.NewEncoder(w).Encode(groupHandler.groupService.GetAllGroupS())
			if err != nil {
				return err
			}
		}
	}
	if r.Method == http.MethodDelete {
		var deleteIds modelsDTO.DeleteDTO
		err := json.NewDecoder(r.Body).Decode(&deleteIds)
		if err != nil {
			return err
		}
		groupHandler.groupService.DeleteGroupS(deleteIds)

		return nil
	}

	return nil
}
