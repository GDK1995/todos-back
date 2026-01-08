package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type UserRepository interface {
	AddUser(user models.User) error
	GetAllUser() []models.User
	GetUserById(userId int) (models.User, error)
	DeleteUser(userId int) error
	GetUserByEmail(email string) (models.User, error)
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
