package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/models"
	"todo/services"
)

type TaskHandler interface {
	TaskHandle(w http.ResponseWriter, r *http.Request) error
}

type taskHandler struct {
	taskService services.TaskService
}

func NewTaskHandler(taskService services.TaskService) TaskHandler {
	return &taskHandler{taskService: taskService}
}

func (taskHandler *taskHandler) TaskHandle(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			return err
		}
		taskHandler.taskService.AddTaskS(task)

		return nil
	}
	if r.Method == http.MethodGet {
		groupIdStr := r.URL.Query().Get("group_id")

		if groupIdStr != "" {
			id, err := strconv.Atoi(groupIdStr)
			if err != nil {
				return err
			}

			tasks := taskHandler.taskService.GetTaskByGroupS(id)

			return json.NewEncoder(w).Encode(tasks)
		}
	}
	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}
		taskHandler.taskService.DeleteTaskS(id)

		return nil
	}

	return nil
}
