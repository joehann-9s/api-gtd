package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gorilla/mux"
	"github.com/joehann-9s/api-gtd/api/models"
	"github.com/joehann-9s/api-gtd/pkg/auth"
	"github.com/joehann-9s/api-gtd/pkg/db"
)

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	UserID := auth.GetIDFromClaims(r)

	err := db.DB.Where("id = ? AND user_id = ?", params["id"], UserID).Preload("SubTasks").First(&task).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "Task not found"})
		return
	}

	response := struct {
		Success bool        `json:"success"`
		Data    models.Task `json:"data"`
	}{
		Success: true,
		Data:    task,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
