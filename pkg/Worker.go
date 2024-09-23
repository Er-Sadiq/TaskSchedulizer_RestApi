package pkg

import (
	"fmt"
	"os/exec"
	"time"
)

// Worker: Execute the command
func Worker(task string, taskTime string) bool {
	fmt.Printf("Worker received Task: %s, Scheduled Time: %s\n", task, taskTime)

	// Parse the scheduled task time
	executionTime, err := time.Parse("2006-01-02 15:04:05", taskTime)
	if err != nil {
		fmt.Printf("Error parsing task time: %v\n", err)
		updateTaskStatus(task, taskTime, false)
		return false
	}

	currentTime := time.Now()
	fmt.Printf("Current time: %s, Execution time: %s\n", currentTime.Format("2006-01-02 15:04:05"), executionTime.Format("2006-01-02 15:04:05"))

	// Check if the scheduled time is in the past
	if currentTime.After(executionTime) {
		fmt.Println("Scheduled time has passed. Executing task immediately.")
	} else {
		// Wait until the scheduled time
		waitDuration := executionTime.Sub(currentTime)
		fmt.Printf("Waiting for %s to execute the task...\n", waitDuration)
		time.Sleep(waitDuration) // Wait until the scheduled time
	}

	// Execute the command
	cmd := exec.Command("bash", "-c", task) // Using bash to execute the command
	output, err := cmd.CombinedOutput()     // Capture output and error
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		fmt.Printf("Output:\n%s\n", output)
		updateTaskStatus(task, taskTime, false)
		return false
	}

	// Print command output
	fmt.Printf("Executing task: %s at %s\nOutput:\n%s\n", task, time.Now().Format("2006-01-02 15:04:05"), output)

	// Update status to success in the database
	updateTaskStatus(task, taskTime, true)
	return true
}

// updateTaskStatus: Update the task status in the database
func updateTaskStatus(task string, time string, success bool) {
	var status int
	if success {
		status = 1 // True
	} else {
		status = 0 // False
	}

	query := "UPDATE TaskTable SET status = ? WHERE task = ? AND time = ?"
	_, err := db.Exec(query, status, task, time)
	if err != nil {
		fmt.Printf("Error updating task status in DB: %v\n", err)
	}
}
