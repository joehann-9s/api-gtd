package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/api/routes"
	"github.com/joehann-9s/api-gtd/pkg/db"
	"github.com/joehann-9s/api-gtd/pkg/utils"
)

func main() {
	utils.LoadEnv()
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	db.DBConnection()

	//migrating DB
	//db.DB.AutoMigrate(models.User{}, models.Task{}, models.SubTask{}, models.Category{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandlers)
	s := r.PathPrefix("/api/v1").Subrouter()

	// Configuración de rutas para tareas
	routes.ConfigureTaskRoutes(s)
	routes.ConfigureAuthRoutes(s)

	http.ListenAndServe(":"+port, r)
}
