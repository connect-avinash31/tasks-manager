// package service expose REST APIs to manage tasks
// It provides functionality to create, retrieve, update, delete and filter tasks.
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"tasks-manager/model"
	"tasks-manager/service"
)

// HandlerCreateTask is the handler function to create a new task
func HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	// defining the task which willl be created
	var newTask model.Task
	// decocogin the requets body into the newTask struct
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Create the task using TaskManager
	if err := service.TaskHandlerInstance().CreateTask(newTask); err != nil {
		log.Println("task creataion failed, with an error :", err)
		http.Error(w, "Failed to create task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// HandleGetTask is the handler function to retrieve a task by its title
func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	if title := r.URL.Query().Get("title"); len(title) > 0 {
		task, err := service.TaskHandlerInstance().GetTask(title)
		if err != nil {
			log.Println("Failed to retrieve task:", err)
			http.Error(w, "Failed to retrieve task: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(task)
	} else if status := r.URL.Query().Get("status"); len(status) > 0 {
		taskList, err := service.TaskHandlerInstance().FilterTasksByStatus(model.TaskStatus(status))
		if err != nil {
			log.Println("Failed to filter tasks by status:", err)
			http.Error(w, "Failed to filter tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(taskList)
	} else {
		taskList, err := service.TaskHandlerInstance().ListAllTasks()
		if err != nil {
			log.Println("Failed to list all tasks:", err)
			http.Error(w, "Failed to list tasks: "+err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(taskList)
	}
}

// HandleUpdateTask is the handler function to update a task's status
func HandleUpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask model.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if err := service.TaskHandlerInstance().UpdateTask(updatedTask); err != nil {
		log.Println("Failed to update task:", err)
		http.Error(w, "Failed to update task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

// HandleDeleteTask is the handler function to delete a task by its title
func HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	if len(title) == 0 {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}
	log.Println("delting the job with title -", title)
	if err := service.TaskHandlerInstance().DeleteTask(title); err != nil {
		log.Println("Failed to delete task:", err)
		http.Error(w, "Failed to delete task: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
