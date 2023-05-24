package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/joehann-9s/api-gtd/api/models"
	"github.com/joehann-9s/api-gtd/pkg/auth"
	"github.com/joehann-9s/api-gtd/pkg/db"
)

func GetllAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	UserID := auth.GetIDFromClaims(r)

	err := db.DB.Where("user_id = ?", UserID).Preload("SubTasks").Find(&tasks).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "None tasks for this user"})
		return
	}
	response := struct {
		Success bool          `json:"success"`
		Data    []models.Task `json:"data"`
	}{
		Success: true,
		Data:    tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
