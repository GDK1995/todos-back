package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type GroupRepository interface {
	AddGroup(group models.Group) (int, error)
	GetAllGroup() []models.Group
	GetGroupsByUser(userId int) []models.Group
	DeleteGroup(groupId int)
}

type groupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) GroupRepository {
	return &groupRepository{db: db}
}

func (groupRepository *groupRepository) AddGroup(group models.Group) (int, error) {
	var id int
	err := groupRepository.db.QueryRow("insert into groups(name) values($1) returning id", group.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (groupRepository *groupRepository) GetAllGroup() []models.Group {
	var groupsList []models.Group
	groups, err := groupRepository.db.Query("select * from groups")
	if err != nil {
		log.Fatal(err)
	}

	for groups.Next() {
		var group models.Group
		errTwo := groups.Scan(&group.ID, &group.Name)
		if errTwo != nil {
			log.Fatal(err)
		}
		groupsList = append(groupsList, group)
	}

	return groupsList
}

func (groupRepository *groupRepository) GetGroupsByUser(userId int) []models.Group {
	var groupList []models.Group
	groups, err := groupRepository.db.Query("select g.id, g.name from groups g join user_groups ug on ug.group_id = g.id where ug.user_id = $1", userId)
	if err != nil {
		log.Fatal(err)
	}

	for groups.Next() {
		var group models.Group
		errTwo := groups.Scan(&group.ID, &group.Name)
		if errTwo != nil {
			log.Fatal(errTwo)
		}
		groupList = append(groupList, group)
	}
	return groupList
}

func (groupRepository *groupRepository) DeleteGroup(groupId int) {
	_, err := groupRepository.db.Exec("delete from groups where id = $1", groupId)
	if err != nil {
		log.Fatal(err)
	}
}
