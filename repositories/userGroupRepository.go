package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type UserGroupRepository interface {
	AddUserToGroup(userGroup models.UserGroup)
	DeleteUserFromGroup(userGroup models.UserGroup)
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

func (userGroupRepository *userGroupRepository) DeleteUserFromGroup(userGroup models.UserGroup) {
	_, err := userGroupRepository.db.Exec("delete from user_groups where user_id = $1 and group_id = $2", userGroup.UserID, userGroup.GroupID)
	if err != nil {
		return
	}
}
