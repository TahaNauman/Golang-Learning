# Task Manager REST API

A simple REST API built with Go's standard library for managing tasks.

Task data is stored in memory and is lost when the server restarts.

## Features

- Create, retrieve, update, and delete tasks.
- Validate task titles and request payloads.
- Return appropriate HTTP status codes.
- Protect shared task data with `sync.Mutex`.
- Unit tests using Go's `net/http/httptest` package.
- No third-party dependencies.

## Tech Stack

- **Language:** Go
- **HTTP:** `net/http`
- **JSON:** `encoding/json`
- **Testing:** `testing`, `net/http/httptest`
- **Concurrency:** `sync.Mutex`

## Prerequisites

- Go 1.22 or later installed on your system.

## How to Run

Clone the repository and enter the project directory:

```bash
git clone https://github.com/TahaNauman/Golang-Learning.git
cd Golang-Learning/task-manager-api
```

Start the server:

```bash
go run .
```

The API runs at `http://localhost:8080` and starts with two sample tasks.

## API Endpoints

| Method | Endpoint | Description | Success |
|---|---|---|---|
| GET | `/` | Check API status | `200 OK` |
| GET | `/tasks` | Retrieve all tasks | `200 OK` |
| POST | `/tasks` | Create a task | `201 Created` |
| PUT | `/tasks/{id}` | Update a task | `200 OK` |
| DELETE | `/tasks/{id}` | Delete a task | `204 No Content` |

### Create a Task

`POST /tasks` with Content-Type `application/json`:

```json
{
  "title": "Learn Go REST APIs"
}
```

The server assigns the ID and initializes `completed` to `false`.

### Update a Task

`PUT /tasks/1` with Content-Type `application/json`:

```json
{
  "title": "Practice Go testing",
  "completed": true
}
```

## Error Handling

- `400 Bad Request` — malformed JSON, an empty task title, or a non-numeric task ID.
- `404 Not Found` — the task ID does not exist.
- `405 Method Not Allowed` — unsupported HTTP method on `/` or `/tasks`.

## Concurrency

The task list is a shared in-memory slice. A `sync.Mutex` guards every read and write, so concurrent requests cannot read or modify partially updated data.

## Running Tests

```bash
go test -v
```

The tests cover all five endpoints, including validation failures, missing tasks, and unsupported methods.

## Current Limitations

- Tasks are stored in memory and reset when the server restarts.
- No database or persistent storage.
- No authentication or authorization.
- Not packaged for production deployment.

## Key Learning Outcomes

- Building REST endpoints with Go's standard library.
- Handling JSON requests and responses.
- Validating input and returning HTTP status codes.
- Writing HTTP handler tests with `httptest`.
- Protecting shared state with a mutex.
