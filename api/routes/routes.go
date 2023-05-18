package routes

import (
	"github.com/gorilla/mux"
)

// Tasks' routes
func ConfigureTaskRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", GetllAllTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", GetTaskByID).Methods("GET")
	r.HandleFunc("/tasks/{state}", GetTaskByState).Methods("GET")
	r.HandleFunc("/tasks/{id}", DeleteTaskByID).Methods("DELETE")
	r.HandleFunc("/tasks/", DeleteAllTasks).Methods("DELETE")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks", UpdateTaskByID).Methods("PATCH")
}

// Users' routes
func ConfigureUserRoutes(r *mux.Router) {

}
