package routes

import (
	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/pkg/auth"
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

// Users auth routes
func ConfigureAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", auth.RegisterUser).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginUser).Methods("POST")

}
