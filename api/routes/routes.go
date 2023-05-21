package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/api/middleware"
	"github.com/joehann-9s/api-gtd/pkg/auth"
)

// Tasks' routes (Protected)
func ConfigureTaskRoutes(r *mux.Router) {
	r.Handle("/tasks", middleware.AuthMiddleware(http.HandlerFunc(CreateTask))).Methods("POST")
	r.Handle("/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(GetTaskByID))).Methods("GET")
	r.Handle("/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(UpdateTaskByID))).Methods("PATCH")
	r.Handle("/tasks/", middleware.AuthMiddleware(http.HandlerFunc(DeleteAllTasks))).Methods("DELETE")

	r.HandleFunc("/tasks", GetllAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{state}", GetTaskByState).Methods("GET")
	r.HandleFunc("/tasks/{id}", DeleteTaskByID).Methods("DELETE")
}

// Users auth routes
func ConfigureAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", auth.RegisterUser).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginUser).Methods("POST")

}
