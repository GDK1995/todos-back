package handlers

import (
	"encoding/json"
	"net/http"
	"todo/models"
	"todo/modelsDTO"
	"todo/services"
)

type UserGroupHandler interface {
	UserGroupHandle(w http.ResponseWriter, r *http.Request)
}

type userGroupHandler struct {
	userGroupService services.UserGroupService
}

func NewUserGroupHandler(userGroupService services.UserGroupService) UserGroupHandler {
	return &userGroupHandler{userGroupService: userGroupService}
}

func (userGroupHandler *userGroupHandler) UserGroupHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var usersGroup modelsDTO.UsersGroupDTO
		err := json.NewDecoder(r.Body).Decode(&usersGroup)
		if err != nil {
			return
		}
		userGroupHandler.userGroupService.AddUserToGroupS(usersGroup)
	}
	if r.Method == http.MethodDelete {
		var userGroup models.UserGroup
		err := json.NewDecoder(r.Body).Decode(&userGroup)
		if err != nil {
			return
		}
		userGroupHandler.userGroupService.DeleteUserFromGroupS(userGroup)
	}
}
