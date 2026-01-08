package modelsDTO

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
