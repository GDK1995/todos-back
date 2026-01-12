package services

import (
	"fmt"
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"
)

type UserGroupService interface {
	AddUserToGroupS(usersGroup modelsDTO.UsersGroupDTO)
	DeleteUserFromGroupS(userGroup models.UserGroup)
}

type userGroupService struct {
	userGroupRepository repositories.UserGroupRepository
}

func NewUserGroupService(userGroupRepository repositories.UserGroupRepository) UserGroupService {
	return &userGroupService{userGroupRepository: userGroupRepository}
}

func (userGroupService *userGroupService) AddUserToGroupS(usersGroup modelsDTO.UsersGroupDTO) {
	ids := usersGroup.UserIDs
	groupId := usersGroup.GroupID
	for _, value := range ids {
		fmt.Println(value)
		userGroup := models.UserGroup{
			UserID:  value,
			GroupID: groupId,
		}
		userGroupService.userGroupRepository.AddUserToGroup(userGroup)
	}
}

func (userGroupService *userGroupService) DeleteUserFromGroupS(userGroup models.UserGroup) {
	userGroupService.userGroupRepository.DeleteUserFromGroup(userGroup)
}
