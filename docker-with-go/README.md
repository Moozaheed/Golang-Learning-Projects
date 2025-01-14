
# Task Management API

This is a simple **Task Management API** built with Go. The API provides endpoints to manage tasks, allowing users to add and fetch tasks. The application is containerized using Docker for easy deployment.

## Features

- **GET `/tasks`**: Fetch all tasks.
- **POST `/add-task`**: Add a new task with a title and description.

---

## Getting Started

Follow the steps below to set up and run the project locally or with Docker.

### Prerequisites

- [Go](https://golang.org/dl/) (for local development)
- [Docker](https://www.docker.com/) (for containerized deployment)

---

## Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/Moozaheed/Golang-Learning-Projects/tree/main/docker-with-go
   cd docker-with-go
   ```

2. Build and run the Go application:
   ```bash
   go build -o main .
   ./main
   ```

3. Access the API on [http://localhost:8080](http://localhost:8080).

---

## Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t docker-with-go:latest .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 -it docker-with-go
   ```

3. Access the API on [http://localhost:8080](http://localhost:8080).

---

## API Endpoints

### 1. GET `/tasks`

**Description**: Retrieve all tasks.

**Example Request**:
```bash
curl http://localhost:8080/tasks
```

**Example Response**:
```json
[
  {
    "id": 1,
    "title": "Complete project",
    "description": "Work on the final project implementation."
  },
  {
    "id": 2,
    "title": "Prepare presentation",
    "description": "Create slides for the upcoming meeting."
  }
]
```

### 2. POST `/add-task`

**Description**: Add a new task.

**Example Request**:
```bash
curl -X POST http://localhost:8080/add-task -H "Content-Type: application/json" -d '{"title":"New Task","description":"Task details"}'
```

**Example Response**:
```json
{
  "id": 3,
  "title": "New Task",
  "description": "Task details"
}
```

---

## Docker Workflow

1. **Build the Docker Image**:
   ```bash
   docker build -t docker-with-go:latest .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -p 8080:8080 -it docker-with-go
   ```




