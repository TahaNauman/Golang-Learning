package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
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

var tasksMutex sync.Mutex

func maxID(tasks []Task) int {
	maxId := 0
	for _, task := range tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}
	return maxId
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Task Manager API!", "status": "running"})
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tasksMutex.Lock()
			tasksCopy := append([]Task(nil), tasks...)
			tasksMutex.Unlock()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasksCopy)
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

			tasksMutex.Lock()
			defer tasksMutex.Unlock()
			newTask.ID = maxID(tasks) + 1
			newTask.Completed = false

			tasks = append(tasks, newTask)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTask)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("PUT /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		var updatedTask Task
		if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(updatedTask.Title) == "" {
			http.Error(w, "Task title is required", http.StatusBadRequest)
			return
		}

		tasksMutex.Lock()

		var result Task
		found := false

		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Title = updatedTask.Title
				tasks[i].Completed = updatedTask.Completed
				result = tasks[i]
				found = true
				break
			}
		}

		tasksMutex.Unlock()

		if !found {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	mux.HandleFunc("DELETE /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		tasksMutex.Lock()
		defer tasksMutex.Unlock()
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		http.Error(w, "Task not found", http.StatusNotFound)
	})

	return mux
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", setupRoutes()))
}
