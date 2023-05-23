package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/joehann-9s/api-gtd/api/models"
	"github.com/joehann-9s/api-gtd/pkg/auth"
	"github.com/joehann-9s/api-gtd/pkg/db"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	task.UserID = auth.GetIDFromClaims(r)

	err := db.DB.Create(&task).Error
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"message": "Failed to create task"})
		return
	}

	response := struct {
		Message string `json:"message"`
		TaskID  uint   `json:"task_id"`
	}{
		Message: "Task created successfully",
		TaskID:  task.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
