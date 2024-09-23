package pkg

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Task string `json:"task"`
	Time string `json:"time"`
}

var db *sql.DB

// SchedulerHandler handles the /taskschedulizer endpoint
func SchedulerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var task Task

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the task into the database
	if err := InsertResponseIntoDB(task.Task, task.Time); err != nil {
		http.Error(w, "Error inserting into database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Sample response to send back to the client
	response := map[string]string{
		"output": "Task scheduled successfully!",
		"task":   task.Task,
		"time":   task.Time,
	}

	// Set the header and encode the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// SetUpDataBase initializes the database connection (to be reused)
func SetUpDataBase() error {
	var err error
	// Only set up the connection if it hasn't been set up yet
	if db == nil {
		db, err = sql.Open("mysql", "root:mysql#888@tcp(127.0.0.1:3306)/SchedulerSchema")
		if err != nil {
			return err
		}
		// Check if the connection
		if err := db.Ping(); err != nil {
			return err
		}
	}
	return nil
}

// InsertResponseIntoDB inserts the task and time into the database
func InsertResponseIntoDB(task string, time string) error {
	if err := SetUpDataBase(); err != nil {
		return fmt.Errorf("error setting up database: %w", err)
	}

	query := "INSERT INTO TaskTable(task, time) VALUES(?, ?)"
	_, err := db.Exec(query, task, time)
	if err != nil {
		return err
	}

	fmt.Println("Task scheduled successfully!")
	return nil
}
