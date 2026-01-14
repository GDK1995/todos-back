package services

import (
	"fmt"
	"todo/models"
	"todo/modelsDTO"
	"todo/repositories"
)

type TaskService interface {
	AddTaskS(task models.Task)
	GetTaskByGroupS(groupId int) []models.Task
	DeleteTaskS(taskId int)
	GetAllTaskS(userId int) []modelsDTO.TaskDTO
	UpdateTaskS(task models.Task)
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

func (taskService *taskService) UpdateTaskS(task models.Task) {
	taskItem := taskService.taskRepository.GetTaskById(task.ID)

	if task.Name != taskItem.Name {
		taskItem.Name = task.Name
	}
	if task.Description != taskItem.Description {
		taskItem.Description = task.Description
	}
	if task.IsDone != taskItem.IsDone {
		taskItem.IsDone = task.IsDone
	}
	if task.Deadline != taskItem.Deadline {
		taskItem.Deadline = task.Deadline
	}
	if task.GroupId != taskItem.GroupId {
		taskItem.GroupId = task.GroupId
	}

	taskService.taskRepository.UpdateTask(task)
}

func (taskService *taskService) GetTaskByGroupS(groupId int) []models.Task {
	return taskService.taskRepository.GetTaskByGroup(groupId)
}

func (taskService *taskService) DeleteTaskS(taskId int) {
	taskService.taskRepository.DeleteTask(taskId)
}

func (taskService *taskService) GetAllTaskS(userId int) []modelsDTO.TaskDTO {
	return taskService.taskRepository.GetAllTask(userId)
}
