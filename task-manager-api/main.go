package main

import (
	"fmt"
	"net/http"
)


func main() {

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Welcome to the Task Manager API!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
})

	http.ListenAndServe(":8080", nil)
}