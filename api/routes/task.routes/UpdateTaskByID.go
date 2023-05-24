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

func UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	params := mux.Vars(r)
	updatedTask.UserID = auth.GetIDFromClaims(r)

	var task models.Task
	err := db.DB.Where("id = ? AND user_id = ?", params["id"], updatedTask.UserID).Preload("SubTasks").First(&task).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "Task not found"})
		return
	}

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.UserID = updatedTask.UserID
	task.Type = updatedTask.Type
	task.State = updatedTask.State
	task.Categories = updatedTask.Categories
	task.ReminderDate = updatedTask.ReminderDate
	task.SubTasks = updatedTask.SubTasks

	err = db.DB.Save(&task).Error
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"message": "Failed to update task"})
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
