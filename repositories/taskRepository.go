package repositories

import (
	"database/sql"
	"log"
	"todo/models"
)

type TaskRepository interface {
	AddTask(task models.Task)
	GetTaskByGroup(groupId int) []models.Task
	DeleteTask(taskId int)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (taskRepository *taskRepository) AddTask(task models.Task) {
	_, err := taskRepository.db.Exec("insert into tasks(name, description, isdone, deadline, group_id) values($1, $2, $3, $4, $5)", task.Name, task.Description, task.IsDone, task.Deadline, task.GroupId)
	if err != nil {
		log.Fatal(err)
	}
}

func (taskRepository *taskRepository) GetTaskByGroup(groupId int) []models.Task {
	var taskList []models.Task
	tasks, err := taskRepository.db.Query("select * from tasks where group_id = $1", groupId)
	if err != nil {
		log.Fatal(err)
	}

	for tasks.Next() {
		var task models.Task
		errTwo := tasks.Scan(&task.ID, &task.Name, &task.Description, &task.IsDone, &task.Deadline, &task.GroupId)
		if errTwo != nil {
			log.Fatal(errTwo)
		}

		taskList = append(taskList, task)
	}

	return taskList
}

func (taskRepository *taskRepository) DeleteTask(taskId int) {
	_, err := taskRepository.db.Exec("delete from tasks where id = $1", taskId)
	if err != nil {
		log.Fatal(err)
	}
}
