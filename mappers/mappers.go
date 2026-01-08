package mappers

import (
	"todo/models"
	"todo/modelsDTO"
)

func MapToUserDTOList(users []models.User) []modelsDTO.UserDTO {
	var userDTOs []modelsDTO.UserDTO
	for i := 0; i < len(users); i++ {
		var userDTO modelsDTO.UserDTO
		userDTO.ID = users[i].ID
		userDTO.Username = users[i].Username
		userDTO.Email = users[i].Email
		userDTOs = append(userDTOs, userDTO)
	}
	return userDTOs
}

func MapToUserDTO(user models.User) modelsDTO.UserDTO {
	var userDTO modelsDTO.UserDTO
	userDTO.ID = user.ID
	userDTO.Username = user.Username
	userDTO.Email = user.Email

	return userDTO
}
