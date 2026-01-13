package main

import (
	"database/sql"
	"log"
	"net/http"
	"todo/handlers"
	"todo/middlewares"
	"todo/repositories"
	"todo/services"

	_ "github.com/lib/pq"
)

import "github.com/rs/cors"

var db *sql.DB

func InitDB() {
	connection := "user=postgres password=Postgresql1995! dbname=todos host=localhost port=5432 sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	errTwo := db.Ping()
	if errTwo != nil {
		log.Fatal(errTwo)
	}
}
func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	InitDB()
	defer CloseDB()

	var userRepository repositories.UserRepository
	userRepo := repositories.NewUserRepository(db)
	userRepository = userRepo

	var groupRepository repositories.GroupRepository
	groupRepo := repositories.NewGroupRepository(db)
	groupRepository = groupRepo

	var userGroupRepostory repositories.UserGroupRepository
	userGroupRepo := repositories.NewUserGroupRepository(db)
	userGroupRepostory = userGroupRepo

	var taskRepository repositories.TaskRepository
	taskRepo := repositories.NewTaskRepository(db)
	taskRepository = taskRepo

	var userService services.UserService
	userServ := services.NewUserService(userRepository)
	userService = userServ

	var groupService services.GroupService
	groupServ := services.NewGroupService(groupRepository)
	groupService = groupServ

	var taskService services.TaskService
	taskServ := services.NewTaskService(taskRepository)
	taskService = taskServ

	var userGroupService services.UserGroupService
	userGroupServ := services.NewUserGroupService(userGroupRepostory)
	userGroupService = userGroupServ

	var authService services.AuthService
	authServ := services.NewAuthService(userRepository)
	authService = authServ

	var userHandler handlers.UserHandler
	userHandle := handlers.NewUserHandler(userService)
	userHandler = userHandle

	var groupHandler handlers.GroupHandler
	groupHandle := handlers.NewGroupHandler(groupService)
	groupHandler = groupHandle

	var userGroupHandler handlers.UserGroupHandler
	userGroupHandle := handlers.NewUserGroupHandler(userGroupService)
	userGroupHandler = userGroupHandle

	var taskHandler handlers.TaskHandler
	taskHandle := handlers.NewTaskHandler(taskService)
	taskHandler = taskHandle

	var authHandler handlers.AuthHandler
	authHandle := handlers.NewAuthHandler(authService)
	authHandler = authHandle

	http.Handle("/user", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(userHandler.UserHandle),
	))
	http.Handle("/group", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(groupHandler.GroupHandle),
	))
	http.HandleFunc("/user-group", userGroupHandler.UserGroupHandle)
	http.Handle("/task", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(taskHandler.TaskHandle),
	))
	http.Handle("/task/all", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(taskHandler.AllTaskHandle),
	))
	http.HandleFunc("/auth", authHandler.RegisterHandle)
	http.HandleFunc("/login", middlewares.ErrorMiddleware(authHandler.LoginHandle))

	mux := http.NewServeMux()
	mux.Handle("/user", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(userHandler.UserHandle),
	))
	mux.Handle("/group", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(groupHandler.GroupHandle),
	))
	mux.HandleFunc("/user-group", userGroupHandler.UserGroupHandle)
	mux.Handle("/task", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(taskHandler.TaskHandle),
	))
	mux.Handle("/task/all", middlewares.AuthMiddleware(
		middlewares.ErrorMiddleware(taskHandler.AllTaskHandle),
	))
	mux.HandleFunc("/auth", authHandler.RegisterHandle)
	mux.HandleFunc("/login", middlewares.ErrorMiddleware(authHandler.LoginHandle))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	server := http.Server{
		Addr:    "localhost:7070",
		Handler: c.Handler(mux),
	}
	errTwo := server.ListenAndServe()
	if errTwo != nil {
		log.Fatal("Server can not start : ", errTwo)
	}
}
