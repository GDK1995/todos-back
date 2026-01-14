package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"todo/auth"
	"todo/middlewares"
	"todo/modelsDTO"
	"todo/services"
)

type AuthHandler interface {
	RegisterHandle(w http.ResponseWriter, r *http.Request)
	LoginHandle(w http.ResponseWriter, r *http.Request) error
	UpdateUser(w http.ResponseWriter, r *http.Request) error
}

type authHandler struct {
	authService services.AuthService
}

type LoginResponse struct {
	Token string            `json:"token"`
	User  modelsDTO.UserDTO `json:"user"`
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (authHandler *authHandler) RegisterHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userDTO modelsDTO.UserAuthDTO
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			return
		}
		errTwo := authHandler.authService.RegisterUser(userDTO)
		if errTwo != nil {
			return
		}
	}
}

func (authHandler *authHandler) LoginHandle(w http.ResponseWriter, r *http.Request) error {
	var login modelsDTO.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		return &middlewares.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid request",
			Err:     err,
		}
	}

	user, token, errTwo := authHandler.authService.Login(login.Email, login.PasswordHash)
	if errTwo != nil {
		if errors.Is(errTwo, auth.ErrInvalidCredentials) {
			return &middlewares.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: "Неверный email или пароль",
				Err:     errTwo,
			}
		}

		return errTwo
	}

	return json.NewEncoder(w).Encode(LoginResponse{
		Token: token,
		User:  user,
	})
}

func (authHandler *authHandler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	var user modelsDTO.UserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return err
	}

	updatedUser, errTwo := authHandler.authService.UpdateUserS(user)
	if errTwo != nil {
		return errTwo
	}

	return json.NewEncoder(w).Encode(updatedUser)
}
