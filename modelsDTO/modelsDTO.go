package modelsDTO

import "time"

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserAuthDTO struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	PasswordPlain string `json:"password"`
}

type LoginDTO struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}
type UsersGroupDTO struct {
	UserIDs []int `json:"user_ids"`
	GroupID int   `json:"group_id"`
}

type DeleteDTO struct {
	DeleteIDs []int `json:"delete_ids"`
}

type AuthDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TaskDTO struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsDone      bool       `json:"isDone"`
	Deadline    *time.Time `json:"deadline"`
	GroupId     int        `json:"group_id"`
	GroupName   string     `json:"group_name"`
}
