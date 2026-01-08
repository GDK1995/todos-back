package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type UserGroupRepository interface {
	AddUserToGroup(userGroup models.UserGroup)
}

type userGroupRepository struct {
	db *sql.DB
}

func NewUserGroupRepository(db *sql.DB) UserGroupRepository {
	return &userGroupRepository{db: db}
}

func (userGroupRepository *userGroupRepository) AddUserToGroup(userGroup models.UserGroup) {
	_, err := userGroupRepository.db.Exec("insert into user_groups(user_id, group_id) values($1, $2) ON CONFLICT (user_id, group_id) DO NOTHING", userGroup.UserID, userGroup.GroupID)
	if err != nil {
		log.Fatal(err)
	}
}
