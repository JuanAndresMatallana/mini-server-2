package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenido a mi mundo")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {

	tasks := []Task{
		{ID: 1, Name: "estudiar go"},
		{ID: 2, Name: "Estudiar documentacion"},
		{ID: 3, Name: "Hacer un proyecto"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

var tasks []Task
var nextID = 1

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/tasks", taskHandler)
	http.ListenAndServe(":8080", nil)
}
