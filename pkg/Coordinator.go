package pkg

import (
	"fmt"
	"log"
)

// Coordinator: Fetch data from DB and assign tasks to workers
func Coordinator() {
	task, time, err := FetchDataFromDB()
	if err != nil {
		log.Printf("Error fetching data from DB: %v", err)
		return
	}

	// Pass task to worker and execute it
	success := AssignToWorker(task, time)
	if success {
		log.Println("Task executed successfully")
	} else {
		log.Println("Task execution failed")
	}
}

// FetchDataFromDB: Fetch task and time from the database
func FetchDataFromDB() (string, string, error) {
	if err := SetUpDataBase(); err != nil {
		return "", "", fmt.Errorf("error setting up database: %w", err)
	}

	rows, err := db.Query("SELECT task, time FROM TaskTable WHERE status = false LIMIT 1")
	if err != nil {
		return "", "", fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var task, time string
	if rows.Next() {
		if err := rows.Scan(&task, &time); err != nil {
			return "", "", fmt.Errorf("failed to scan row: %w", err)
		}
		log.Printf("Fetched Task: %s, Time: %s\n", task, time)
		return task, time, nil
	}

	log.Println("No tasks available in the database.")
	return "", "", fmt.Errorf("no tasks available")
}

// AssignToWorker: Send task and time to the worker and update status in DB
func AssignToWorker(task string, time string) bool {
	success := Worker(task, time)

	// Update task status in the database based on worker execution result
	var status int
	if success {
		status = 1 // True
	} else {
		status = 0 // False
	}

	query := "UPDATE TaskTable SET status = ? WHERE task = ? AND time = ?"
	_, err := db.Exec(query, status, task, time)
	if err != nil {
		log.Printf("Error updating task status: %v", err)
		return false
	}

	return success
}

// CheckStatus: Check the status of a task by its ID
func CheckStatus(taskID int) (bool, error) {
	var status int
	err := db.QueryRow("SELECT status FROM TaskTable WHERE id = ?", taskID).Scan(&status)
	if err != nil {
		log.Printf("Error fetching task status: %v", err)
		return false, err
	}

	return status == 1, nil
}
