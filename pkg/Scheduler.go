package pkg

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Task string `json:"task"`
	Time string `json:"time"`
}

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

	//printTask(task)

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

// // printTask prints the details of a given task
// func printTask(t Task) {
// 	fmt.Println("Task:", t.Task)
// 	fmt.Println("Scheduled Time:", t.Time)
// }
