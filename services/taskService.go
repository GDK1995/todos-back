package services

import (
	"todo/models"
	"todo/repositories"
)

type TaskService interface {
	AddTaskS(task models.Task)
	GetTaskByGroupS(groupId int) []models.Task
	DeleteTaskS(taskId int)
}

type taskService struct {
	taskRepository repositories.TaskRepository
}

func NewTaskService(taskRepository repositories.TaskRepository) TaskService {
	return &taskService{taskRepository: taskRepository}
}

func (taskService *taskService) AddTaskS(task models.Task) {
	taskService.taskRepository.AddTask(task)
}

func (taskService *taskService) GetTaskByGroupS(groupId int) []models.Task {
	return taskService.taskRepository.GetTaskByGroup(groupId)
}

func (taskService *taskService) DeleteTaskS(taskId int) {
	taskService.taskRepository.DeleteTask(taskId)
}
