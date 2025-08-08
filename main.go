package main

import (
	"log"
	"net/http"
	"tasks-manager/config"
	"tasks-manager/handler"
	"tasks-manager/service"
)

func main() {
	// Initialize the configuration first
	cfg := config.LoadConfig()
	// Now intilaize the task manager service
	if err := service.NewTaskManager(cfg); err != nil {
		log.Fatalf("Failed to initialize task manager: %v", err)
	}

	// Start the HTTP server with the task manager handler
	http.HandleFunc("/tasks", handler.HandleCreateTask)
	http.HandleFunc("/tasks/get", handler.HandleGetTask)
	http.HandleFunc("/tasks/update", handler.HandleUpdateTask)
	http.HandleFunc("/tasks/delete", handler.HandleDeleteTask)
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
