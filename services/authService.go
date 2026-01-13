package services

import (
	"todo/auth"
	"todo/mappers"
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(userDTO modelsDTO.UserAuthDTO) error
	Login(email, password string) (modelsDTO.UserDTO, string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{userRepository: userRepository}
}

func (authService *authService) RegisterUser(userDTO modelsDTO.UserAuthDTO) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(userDTO.PasswordPlain),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	userItem := models.User{
		Username:     userDTO.Username,
		Email:        userDTO.Email,
		PasswordHash: string(hash),
	}

	errTwo := authService.userRepository.AddUser(userItem)
	if errTwo != nil {
		return errTwo
	}

	return nil
}

func (authService *authService) Login(email, password string) (modelsDTO.UserDTO, string, error) {
	user, err := authService.userRepository.GetUserByEmail(email)

	if err != nil {
		return modelsDTO.UserDTO{}, "", auth.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return modelsDTO.UserDTO{}, "", auth.ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return modelsDTO.UserDTO{}, "", err
	}

	userDTO := mappers.MapToUserDTO(user)

	return userDTO, token, nil
}
