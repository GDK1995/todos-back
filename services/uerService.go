package services

import (
	"todo/mappers"
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"
)

type UserService interface {
	GetAllUserS() []modelsDTO.UserDTO
	GetUserByIdS(userId int) (models.User, error)
	DeleteUserS(deleteIds modelsDTO.DeleteDTO) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (userService *userService) GetAllUserS() []modelsDTO.UserDTO {
	users := userService.userRepository.GetAllUser()
	userDTOs := mappers.MapToUserDTOList(users)
	return userDTOs
}

func (userService *userService) GetUserByIdS(userId int) (models.User, error) {
	id, err := userService.userRepository.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}

	return id, nil
}

func (userService *userService) DeleteUserS(deleteIds modelsDTO.DeleteDTO) error {
	for _, id := range deleteIds.DeleteIDs {
		err := userService.userRepository.DeleteUser(id)
		if err != nil {
			return err
		}
	}

	return nil
}
