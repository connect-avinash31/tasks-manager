// package strorage is the package that provides database operations for the task manager system.
package storage

import (
	"database/sql"
	"tasks-manager/config"
	"tasks-manager/model"

	_ "github.com/lib/pq"
)

// InitalizeDatabase initializes the database connection using the provided configuration.
func InitializeDatabase(cfg *config.Config) (*sql.DB, error) {
	// Connect to the database using the provided URL
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is established
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// Initialize the database schema if necessary
	if err := IntilizeSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

// IntilizeSchema initializes the database schema by creating necessary tables.
func IntilizeSchema(db *sql.DB) error {

	// Create the tasks table if it does not exist
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		title TEXT PRIMARY KEY,
		description TEXT,
		due_date TEXT,
		status TEXT 
	)`
	_, err := db.Exec(query)
	return err
}

// Store task will use the database connection to store a task in the database.
func StoreTask(db *sql.DB, task model.Task) error {
	query := `INSERT INTO tasks (title, description, due_date, status) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, task.Title, task.Description, task.DueDate, task.Status)
	return err
}

// GetTask retrieves a task by its title from the database.
func GetTask(db *sql.DB, title string) (*model.Task, error) {
	query := `SELECT title, description, due_date, status FROM tasks WHERE title = $1`
	row := db.QueryRow(query, title)

	// Create a Task to hold the result
	task := &model.Task{}
	err := row.Scan(&task.Title, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// FilterTasksByStatus retrieves tasks from the database that match the given status.
func FilterTasksByStatus(db *sql.DB, status model.TaskStatus) ([]model.Task, error) {
	query := `SELECT title, description, due_date, status FROM tasks WHERE status = $1`
	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetAllTasks retrieves all tasks from the database.
func GetAllTasks(db *sql.DB) ([]model.Task, error) {
	query := `SELECT title, description, due_date, status FROM tasks`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Title, &task.Description, &task.DueDate, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// UpdateTask will update whole task details in the database.
func UpdateTask(db *sql.DB, task model.Task) error {
	query := `UPDATE tasks SET description = $1, due_date = $2, status = $3 WHERE title = $4`
	_, err := db.Exec(query, task.Description, task.DueDate, task.Status, task.Title)
	return err
}

// DeleteTask removes a task from the database by its title.
func DeleteTask(db *sql.DB, title string) error {
	query := `DELETE FROM tasks WHERE title = $1`
	_, err := db.Exec(query, title)
	return err
}
