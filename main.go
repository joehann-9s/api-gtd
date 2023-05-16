package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/api/routes"
	"github.com/joehann-9s/api-gtd/pkg/db"
	"github.com/joehann-9s/api-gtd/pkg/utils"
)

func main() {
	utils.LoadEnv()
	db.DBConnection()
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandlers)

	http.ListenAndServe(":7000", r)
}