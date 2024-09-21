package main

import (
	"fmt"
	"net/http"

	"github.com/Er-Sadiq/TaskSchedulizer-api/pkg"
)

// function to start the server
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

	http.HandleFunc("/taskschedulizer", pkg.SchedulerHandler)

	go startServer()

	select {}

}
