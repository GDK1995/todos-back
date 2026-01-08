package services

import (
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"
)

type GroupService interface {
	AddGroupS(group models.Group) (int, error)
	GetAllGroupS() []models.Group
	GetGroupsByUserS(userId int) []models.Group
	DeleteGroupS(groupIds modelsDTO.DeleteDTO)
}

type groupService struct {
	groupRepository repositories.GroupRepository
}

func NewGroupService(groupRepository repositories.GroupRepository) GroupService {
	return &groupService{groupRepository: groupRepository}
}

func (groupService *groupService) AddGroupS(group models.Group) (int, error) {
	addedId, err := groupService.groupRepository.AddGroup(group)
	if err != nil {
		return 0, err
	}

	return addedId, nil
}

func (groupService *groupService) GetAllGroupS() []models.Group {
	return groupService.groupRepository.GetAllGroup()
}

func (groupService *groupService) GetGroupsByUserS(userId int) []models.Group {
	return groupService.groupRepository.GetGroupsByUser(userId)
}

func (groupService *groupService) DeleteGroupS(groupIds modelsDTO.DeleteDTO) {
	for _, value := range groupIds.DeleteIDs {
		groupService.groupRepository.DeleteGroup(value)
	}
}
