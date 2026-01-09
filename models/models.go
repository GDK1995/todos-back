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
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsDone      bool       `json:"isDone"`
	Deadline    *time.Time `json:"deadline"`
	GroupId     int        `json:"group_id"`
}

type TaskUser struct {
	TaskID int
	UserID int
}
