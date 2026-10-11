package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{ID: 1, Title: "Task 1", Completed: false},
	{ID: 2, Title: "Task 2", Completed: true},
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Task Manager API!", "status": "running"})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks)
		} else if r.Method == http.MethodPost {
			var newTask Task
			if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			if strings.TrimSpace(newTask.Title) == "" {
				http.Error(w, "Task title is required", http.StatusBadRequest)
				return
			}

			newTask.ID = len(tasks) + 1
			newTask.Completed = false

			tasks = append(tasks, newTask)
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newTask)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
