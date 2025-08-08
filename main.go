package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks []Task
var nextID = 1

//una funcion para generar un handler

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenido a mi mundo")
}

func taskHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(tasks)
		return

	}
	if r.Method == http.MethodPost {

		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Error al leer una tarea", http.StatusBadRequest)
			return
		}

		newTask.ID = nextID
		nextID++
		tasks = append(tasks, newTask)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
		return
	}

	http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)

}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	for i, task := range tasks {

		if task.ID == id {

			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return

		}

	}
	http.Error(w, "Tarea no encontrada", http.StatusNotFound)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "ID invalido", http.StatusNotFound)
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.Error(w, "Tarea no encontrada", http.StatusBadRequest)

}

func main() {
	tasks = append(tasks, Task{ID: 1, Name: "estudiar go"})
	tasks = append(tasks, Task{ID: 2, Name: "Estudiar documentacion"})
	tasks = append(tasks, Task{ID: 3, Name: "Hacer un proyecto"})
	nextID = 4

	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	r.HandleFunc("/tasks", taskHandler).Methods("GET", "POST")
	r.HandleFunc("/tasks/{id}", deleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", getTaskHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}
