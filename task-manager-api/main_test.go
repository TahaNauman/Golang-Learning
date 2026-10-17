package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func resetTasks() {

	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tasks = []Task{
		{ID: 1, Title: "Task 1", Completed: false},
		{ID: 2, Title: "Task 2", Completed: true},
	}
}

func TestHealthEndpoint(t *testing.T) {
	resetTasks()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",
			http.StatusOK, rec.Code)
	}

	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json content type")
	}
}

func TestGetTasks(t *testing.T) {
	resetTasks()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",
			http.StatusOK, rec.Code)
	}

	var got []Task
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(got) == 0 {
		t.Error("expected at least one task")
	}
}

func TestCreateTask(t *testing.T) {
	resetTasks()
	body := strings.NewReader(`{"title":"Learn Go testing"}`)

	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d",
			http.StatusCreated, rec.Code)
	}

	var got Task
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Title != "Learn Go testing" {
		t.Errorf("expected title %q, got %q",
			"Learn Go testing", got.Title)
	}

	if got.ID <= 0 {
		t.Errorf("expected a positive task ID, got %d", got.ID)
	}

	if got.Completed {
		t.Error("expected new task to be incomplete")
	}
}

func TestCreateTaskWithEmptyTitle(t *testing.T) {
	resetTasks()
	body := strings.NewReader(`{"title":"   "}`)

	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d",
			http.StatusBadRequest, rec.Code)
	}
}

func TestCreateTaskWithInvalidJSON(t *testing.T) {
	resetTasks()
	body := strings.NewReader(`{"title":`)

	req := httptest.NewRequest(http.MethodPost, "/tasks", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d",
			http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	resetTasks()
	body := strings.NewReader(
		`{"title":"Updated task","completed":true}`,
	)

	req := httptest.NewRequest(http.MethodPut, "/tasks/1", body)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d",
			http.StatusOK, rec.Code)
	}

	var got Task
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}

	if got.Title != "Updated task" {
		t.Errorf("expected updated title, got %q", got.Title)
	}

	if !got.Completed {
		t.Error("expected task to be completed")
	}
}

func TestDeleteTask(t *testing.T) {
	resetTasks()
	req := httptest.NewRequest(http.MethodDelete, "/tasks/1", nil)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d",
			http.StatusNoContent, rec.Code)
	}
}

func TestUpdateNonexistentTask(t *testing.T) {
	resetTasks()
	req := httptest.NewRequest(http.MethodPut, "/tasks/999",
		strings.NewReader(`{"title":"Missing task","completed":false}`))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d",
			http.StatusNotFound, rec.Code)
	}
}

func TestUpdateTaskWithInvalidID(t *testing.T) {
	resetTasks()

	req := httptest.NewRequest(
		http.MethodPut,
		"/tasks/abc",
		strings.NewReader(`{"title":"Updated task","completed":true}`),
	)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d",
			http.StatusBadRequest, rec.Code)
	}
}

func TestTasksMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPatch, "/tasks", nil)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d",
			http.StatusMethodNotAllowed, rec.Code)
	}
}

func TestDeleteNonexistentTask(t *testing.T) {
	resetTasks()

	req := httptest.NewRequest(
		http.MethodDelete,
		"/tasks/999",
		nil,
	)
	rec := httptest.NewRecorder()

	setupRoutes().ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d",
			http.StatusNotFound, rec.Code)
	}
}
