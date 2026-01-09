package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"todo/modelsDTO"
	"todo/services"
)

type UserHandler interface {
	UserHandle(w http.ResponseWriter, r *http.Request) error
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService: userService}
}

func (userHandler *userHandler) UserHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodDelete {
		var deleteIds modelsDTO.DeleteDTO
		err := json.NewDecoder(r.Body).Decode(userHandler.userService.DeleteUserS(deleteIds))
		if err != nil {
			return err
		}
	}

	if r.Method == http.MethodGet {
		userIdStr := r.URL.Query().Get("user_id")
		groupIdStr := r.URL.Query().Get("group_id")
		gIdStr := r.URL.Query().Get("id")

		if userIdStr != "" && groupIdStr == "" && gIdStr == "" {
			id, err := strconv.Atoi(userIdStr)
			if err != nil {
				http.Error(w, "invalid user id", http.StatusBadRequest)
				return err
			}

			user, errTwo := userHandler.userService.GetUserByIdS(id)
			if errTwo != nil {
				if errors.Is(errTwo, sql.ErrNoRows) {
					http.Error(w, "user not found", http.StatusNotFound)
					return errTwo
				}
				http.Error(w, "internal error", http.StatusInternalServerError)
				return errTwo
			}
			if errThree := json.NewEncoder(w).Encode(user); errThree != nil {
				http.Error(w, "encoding error", http.StatusInternalServerError)
				return errThree
			}
		} else if userIdStr == "" && groupIdStr != "" && gIdStr == "" {
			id, err := strconv.Atoi(groupIdStr)
			if err != nil {
				return err
			}

			users, errTwo := userHandler.userService.GetUsersByGroupIDS(id)
			if errTwo != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
				return errTwo
			}

			errThree := json.NewEncoder(w).Encode(users)
			if errThree != nil {
				http.Error(w, "encoding error", http.StatusInternalServerError)
				return errThree
			}
		} else if userIdStr == "" && groupIdStr == "" && gIdStr != "" {
			fmt.Println("id", gIdStr)
			id, err := strconv.Atoi(gIdStr)
			if err != nil {
				return err
			}

			usersId, errTwo := userHandler.userService.GetUserIdsByGroupIdS(id)
			if errTwo != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
				return errTwo
			}

			errThree := json.NewEncoder(w).Encode(usersId)
			if errThree != nil {
				http.Error(w, "encoding error", http.StatusInternalServerError)
				return errThree
			}
		} else {
			err := json.NewEncoder(w).Encode(userHandler.userService.GetAllUserS())
			if err != nil {
				http.Error(w, "invalid group id", http.StatusBadRequest)
				return err
			}
		}
	}

	return nil
}
