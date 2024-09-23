
# TaskSchedulizer - REST API Task Scheduler

TaskSchedulizer is a RESTful API service that allows users to schedule system tasks to run at specified times. It is built with Go for the backend and utilizes a MySQL (or similar) database to store tasks and their execution statuses. Tasks are executed as shell commands at the scheduled times using `exec.Command` in Go.

## Key Features

- **Task Scheduling**: Users can schedule tasks by sending a task command and the desired execution time via a REST API.
- **Task Execution**: The system automatically executes scheduled tasks at the specified time.
- **Task Status Management**: Tracks the status (success or failure) of each task and updates the database after execution.
- **Status Checking**: Users can check the execution status of previously scheduled tasks.
- **Concurrency Management**: Capable of handling multiple tasks with their respective scheduled times and executing them concurrently.
- **Database Integration**: Uses an SQL database to persist tasks and track their status across restarts.


## Core Components

- **Coordinator**: Manages task fetching from the database and assigns them to the worker for execution.
- **Worker**: Executes the system commands at the scheduled times and updates the task status in the database.
- **Database Interaction**: Performs CRUD operations to manage tasks and update their status after execution.

## API Endpoints

- **POST /taskschedulizer**: Schedule a new task by providing the shell command to execute and the execution time.
- **GET /status**: Retrieve the status of a scheduled task based on the task ID.

## Technologies Used

- **Backend**: Golang
- **Database**: MySQL or SQLite (configurable)
- **Containerization**: Docker for easy deployment and environment management
- **Shell Command Execution**: `exec.Command` to run shell commands on the server

## How It Works

1. A user schedules a task via a REST API call, providing the shell command and the desired execution time.
2. The task is saved in the database with a pending status.
3. At the scheduled time, the Worker retrieves the task from the database, executes the command, and updates the task's status (success or failure).
4. Users can check the status of previously scheduled tasks to see if they ran successfully.

## Installation



```bash
Clone the repository: git clone https://github.com/Er-Sadiq/TaskSchedulizer_RestApi.git

Build the Docker image: docker build -t task_schedulizer .

Run the container: docker run -p 8081:8081 task_schedulizer
```

```bash
# Schedule a new task
curl -X POST http://localhost:8081/taskschedulizer -d '{"task":"whoami", "time":"2024-09-22 17:00:00"}' -H "Content-Type: application/json"

# Check the status of a task
curl -X GET "http://localhost:8081/status?id=1"

```
    
## Demo
 -- Hitting API endpoint
![2024-09-22_17-15_1](https://github.com/user-attachments/assets/d69f9be8-c4ed-4f12-b05e-d244c6e1d8ce)

-- Task getting Scheduled 
![2024-09-22_17-15](https://github.com/user-attachments/assets/5e850052-f154-4068-b5a5-2ebe7c7fbb5d)
