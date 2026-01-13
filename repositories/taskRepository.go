package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"todo/models"
	"todo/modelsDTO"
)

type TaskRepository interface {
	AddTask(task models.Task)
	GetTaskByGroup(groupId int) []models.Task
	DeleteTask(taskId int)
	GetAllTask(userId int) []modelsDTO.TaskDTO
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

func (taskRepository *taskRepository) GetAllTask(userId int) []modelsDTO.TaskDTO {
	fmt.Println(userId)
	tasks, err := taskRepository.db.Query("select t.id, t.name, t.description, t.isdone, t.deadline, t.group_id, g.name from tasks t join groups g on g.id = t.group_id join user_groups ug on ug.group_id = t.group_id where user_id = $1", userId)
	if err != nil {
		return []modelsDTO.TaskDTO{}
	}

	fmt.Println(tasks)

	taskList := make([]modelsDTO.TaskDTO, 0)

	for tasks.Next() {
		var task modelsDTO.TaskDTO
		errTwo := tasks.Scan(&task.ID, &task.Name, &task.Description, &task.IsDone, &task.Deadline, &task.GroupId, &task.GroupName)
		if errTwo != nil {
			return []modelsDTO.TaskDTO{}
		}

		fmt.Println(task)
		taskList = append(taskList, task)
	}

	return taskList
}
