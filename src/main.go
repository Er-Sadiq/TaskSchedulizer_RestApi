package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Er-Sadiq/TaskSchedulizer-api/pkg"
)

func startServer() {
	fmt.Println("TaskSchedulizer's Server Is Up & Running..")
	fmt.Println("Listening at port :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error Loading The Server")
	}
}

func main() {
	fmt.Println("TaskSchedulizer Initializing...")

	// Setup database connection
	if err := pkg.SetUpDataBase(); err != nil {
		fmt.Printf("Error setting up database: %v\n", err)
		return
	}

	http.HandleFunc("/taskschedulizer", pkg.SchedulerHandler)
	http.HandleFunc("/status", pkg.StatusHandler) // Use the correct handler function

	go startServer()

	// Periodically check for tasks to assign to workers
	go func() {
		for {
			time.Sleep(10 * time.Second)
			pkg.Coordinator()
		}
	}()

	select {}
}
