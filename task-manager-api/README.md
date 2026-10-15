# Task Manager REST API

A simple REST API built with Go's standard library to manage tasks. This project explores HTTP handlers, JSON, CRUD operations, input validation, concurrency protection, and API testing.

## Features

- Create, retrieve, update, and delete tasks
- JSON request and response handling
- Input validation for task titles and request payloads
- Appropriate HTTP status codes
- Mutex-based protection for shared task data
- Unit tests using Go's `net/http/httptest` package
- No third-party dependencies

## Tech Stack

- **Language:** Go
- **HTTP:** `net/http`
- **JSON:** `encoding/json`
- **Testing:** `testing`, `net/http/httptest`
- **Concurrency:** `sync.Mutex`

## Getting Started

### Prerequisites

- Go installed on your machine

### Run the API

Clone the repository and enter the project directory:

```bash
git clone https://github.com/TahaNauman/Golang-Learning.git
cd task-manager-api
```

Start the server:

```bash
go run .
```

The API runs at:

`http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/` | Check API status |
| GET | `/tasks` | Retrieve all tasks |
| POST | `/tasks` | Create a task |
| PUT | `/tasks/{id}` | Update a task |
| DELETE | `/tasks/{id}` | Delete a task |

### 1. Health Check

**Request:** `GET /`

Returns a welcome message and API status.

### 2. Get All Tasks

**Request:** `GET /tasks`

Returns the list of tasks as JSON.

### 3. Create a Task

**Request:** `POST /tasks`

Content-Type: `application/json`

Example request body:

```json
{
  "title": "Learn Go REST APIs"
}
```

A successful request returns `201 Created` and the newly created task. The server assigns the ID and initializes `completed` to `false`.

### 4. Update a Task

**Request:** `PUT /tasks/1`

Content-Type: `application/json`

Example request body:

```json
{
  "title": "Practice Go testing",
  "completed": true
}
```

A successful request returns the updated task.

### 5. Delete a Task

**Request:** `DELETE /tasks/1`

A successful deletion returns `204 No Content`.

## Error Handling

The API handles common errors, including:

- `400 Bad Request` — malformed JSON, invalid task ID, or an empty task title
- `404 Not Found` — task ID does not exist
- `405 Method Not Allowed` — unsupported HTTP method for the root or task collection endpoint

## Running Tests

Run the test suite:

```bash
go test -v
```

Run the tests repeatedly in shuffled order:

```bash
go test -shuffle=on -count=5
```

## Current Limitations

- Tasks are stored in memory and reset when the server restarts.
- There is no database or persistent storage.
- Authentication and authorization are not implemented.
- The API has not yet been packaged for production deployment.

## What I Learned

- Building REST endpoints with Go's standard library
- Handling JSON requests and responses
- Validating user input and returning HTTP status codes
- Writing HTTP handler tests with `httptest`
- Protecting shared state with a mutex
