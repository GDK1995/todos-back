package models

import (
	"time"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserGroup struct {
	UserID  int
	GroupID int
}

type Task struct {
	ID          int
	Name        string
	Description string
	IsDone      bool
	Deadline    *time.Time
	GroupId     int
}

type TaskUser struct {
	TaskID int
	UserID int
}
