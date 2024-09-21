# TaskSchedulizer

TaskSchedulizer is an efficient task scheduler written in Go. It allows clients to submit tasks along with their scheduled times, handling the scheduling process seamlessly. This project leverages MySQL for data storage and Docker for containerization.

## Features

- **Task Scheduling**: Accepts tasks and their scheduled times from clients.
- **Efficient Workflow**:
  - **Entry Point**: Interface for clients to submit tasks.
  - **Scheduler**: Manages the scheduling of tasks.
  - **Coordinator**: Oversees workflow and routes tasks.
  - **Worker**: Executes scheduled tasks.
  - **Task Executioner**: Handles the execution of tasks.
- **Database Integration**: Utilizes MySQL for data storage.
- **Containerization**: Built with Docker for easy deployment.

## How It Works

1. **Client Submission**: Clients send tasks and execution times to the entry point.
2. **Task Scheduling**: The Scheduler organizes tasks based on timing.
3. **Coordination**: The Coordinator manages task flow.
4. **Execution**: Workers execute the tasks via the Task Executioner.
5. **Efficiency**: Optimizes resource usage and task handling.

## Technologies Used

- **Programming Language**: Go
- **Database**: MySQL
- **Containerization**: Docker

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/TaskSchedulizer.git

2. Clone the repository:
   ```bash
   cd TaskSchedulizer
   docker build -t taskschedulizer .
   docker run -p 8081:8081 taskschedulizer

