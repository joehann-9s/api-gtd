package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/api/middleware"
	task_routes "github.com/joehann-9s/api-gtd/api/routes/task.routes"
	"github.com/joehann-9s/api-gtd/pkg/auth"
)

// Tasks' routes (Protected)
func ConfigureTaskRoutes(r *mux.Router) {
	r.Handle("/tasks", middleware.AuthMiddleware(http.HandlerFunc(task_routes.CreateTask))).Methods("POST")
	r.Handle("/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(task_routes.GetTaskByID))).Methods("GET")
	r.Handle("/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(task_routes.UpdateTaskByID))).Methods("PUT")
	r.Handle("/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(task_routes.DeleteTaskByID))).Methods("DELETE")
	r.Handle("/tasks", middleware.AuthMiddleware(http.HandlerFunc(task_routes.GetllAllTasks))).Methods("GET")
}

// Users auth routes
func ConfigureAuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", auth.RegisterUser).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginUser).Methods("POST")
}
