package handlers

import (
	"encoding/json"
	"net/http"
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
		json.NewDecoder(r.Body).Decode(&usersGroup)
		userGroupHandler.userGroupService.AddUserToGroupS(usersGroup)
	}
}
