package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type UserRepository interface {
	AddUser(user models.User) error
	UpdateUser(user models.User) error
	GetAllUser() []models.User
	GetUserById(userId int) (models.User, error)
	DeleteUser(userId int) error
	GetUserByEmail(email string) (models.User, error)
	GetUsersByGroupID(groupId int) ([]models.User, error)
	GetUserIdsByGroupID(groupId int) ([]int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) AddUser(user models.User) error {
	_, err := userRepository.db.Exec("insert into users(username, email, password_hash) values($1, $2, $3)", user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *userRepository) UpdateUser(user models.User) error {
	_, err := userRepository.db.Exec("update users set username = $1, password_hash = $2 where id = $3", user.Username, user.PasswordHash, user.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func (userRepository *userRepository) GetAllUser() []models.User {
	var userListItems []models.User
	userList, err := userRepository.db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	for userList.Next() {
		var user models.User
		errTwo := userList.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
		if errTwo != nil {
			log.Fatal(errTwo)
		}
		userListItems = append(userListItems, user)
	}

	return userListItems
}

func (userRepository *userRepository) GetUserById(userId int) (models.User, error) {
	userRow := userRepository.db.QueryRow("select * from users where id = $1", userId)

	var user models.User
	err := userRow.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return user, nil
}

func (userRepository *userRepository) GetUserByEmail(email string) (models.User, error) {
	userRow := userRepository.db.QueryRow("select * from users where email = $1", email)

	var user models.User
	err := userRow.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, err
		}
		return models.User{}, err
	}
	return user, nil
}

func (userRepository *userRepository) DeleteUser(userId int) error {
	_, err := userRepository.db.Exec("delete from users where id = $1", userId)
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *userRepository) GetUsersByGroupID(groupId int) ([]models.User, error) {
	users, err := userRepository.db.Query("select u.id, u.username, u.email from users u join user_groups ug on ug.user_id = u.id where ug.group_id = $1", groupId)
	if err != nil {
		return nil, err
	}

	userList := make([]models.User, 0)

	for users.Next() {
		var user models.User
		errTwo := users.Scan(&user.ID, &user.Username, &user.Email)
		if errTwo != nil {
			return nil, errTwo
		}

		userList = append(userList, user)
	}

	return userList, nil
}

func (userRepository *userRepository) GetUserIdsByGroupID(groupId int) ([]int, error) {
	userIds, err := userRepository.db.Query("select u.id from users u join user_groups ug on ug.user_id = u.id where ug.group_id = $1", groupId)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)

	for userIds.Next() {
		var id int
		errTwo := userIds.Scan(&id)
		if errTwo != nil {
			return nil, errTwo
		}
		ids = append(ids, id)
	}

	return ids, nil
}
