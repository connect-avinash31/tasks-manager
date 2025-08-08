package service

import (
	"database/sql"
	"tasks-manager/config"
	"tasks-manager/model"
	"tasks-manager/storage"
)

// define a task handler interface which will be implemented by the task manager
type TaskHandler interface {
	// CreateTask creates a new task with the given details
	CreateTask(task model.Task) error
	// GetTask retrieves a task by its title
	GetTask(title string) (*model.Task, error)
	// Filter tasks by status
	FilterTasksByStatus(status model.TaskStatus) ([]model.Task, error)
	// ListAllTasks lists all tasks with their details
	ListAllTasks() ([]model.Task, error)
	// UpdateTask updates the status of a task
	UpdateTask(task model.Task) error
	// DeleteTask deletes a task by its title
	DeleteTask(title string) error
}

// TaskManager i sthe implenter of TaskHandler interface
type TaskManager struct {
	db *sql.DB // database connection for storing tasks
}

// taskManager is a singleton instance of TaskManager, so only single instace presents
var TaskMgr TaskHandler

// NewTaskManager initializes a new TaskManager instance
func NewTaskManager(config *config.Config) error {
	// intilaize the Database connection
	db, err := storage.InitializeDatabase(config)
	if err != nil {
		return err
	}
	if TaskMgr == nil {
		TaskMgr = TaskManager{
			db: db,
		}
	}
	return nil
}

// Fetch the singleton instance of TaskManager
func TaskHandlerInstance() TaskHandler {
	if TaskMgr == nil {
		panic("TaskManager is not initialized. Call NewTaskManager first.")
	}
	return TaskMgr
}

// CreateTask creates a new task with the given details
func (tm TaskManager) CreateTask(task model.Task) error {
	// we will call the storage function to store the task in the database
	return storage.StoreTask(tm.db, task)
}

// GetTask retrieves a task by its title
func (tm TaskManager) GetTask(title string) (*model.Task, error) {
	// get the task using title of the task
	return storage.GetTask(tm.db, title)
}

// FilterTasksByStatus retrieves tasks by their status
func (tm TaskManager) FilterTasksByStatus(status model.TaskStatus) ([]model.Task, error) {
	// we will use task status to fiulter out tasks
	return storage.FilterTasksByStatus(tm.db, status)
}

// ListAllTasks retrieves all tasks from the database
func (tm TaskManager) ListAllTasks() ([]model.Task, error) {
	// get the list of all tasks from the database
	return storage.GetAllTasks(tm.db)
}

// UpdateTask updates the status of a task
func (tm TaskManager) UpdateTask(task model.Task) error {
	// first we will check if the task exists
	_, err := storage.GetTask(tm.db, task.Title)
	if err != nil {
		return err
	}
	// now as task presents we will update the status
	return storage.UpdateTask(tm.db, task)
}

// DeleteTask deletes a task by its title
func (tm TaskManager) DeleteTask(title string) error {
	// we will call the storage function to delete the task from the database
	return storage.DeleteTask(tm.db, title)
}
