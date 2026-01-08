package services

import (
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"
)

type UserGroupService interface {
	AddUserToGroupS(usersGroup modelsDTO.UsersGroupDTO)
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
		userGroup := models.UserGroup{
			UserID:  value,
			GroupID: groupId,
		}
		userGroupService.userGroupRepository.AddUserToGroup(userGroup)
	}
}
