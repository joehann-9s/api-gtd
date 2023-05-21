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
		Message models.Task `json:"message"`
	}{
		Message: task,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

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
		Message []models.Task `json:"message"`
	}{
		Message: tasks,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

/*******************/
func GetTaskByState(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Viewing task by state"))

}

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
		Message models.Task `json:"message"`
	}{
		Message: task,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	UserID := auth.GetIDFromClaims(r)

	err := db.DB.Where("id = ? AND user_id = ?", params["id"], UserID).Preload("SubTasks").First(&task).Error
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"message": "Task not found"})
		return
	}

	db.DB.Unscoped().Delete(&task)

	w.WriteHeader(http.StatusNoContent)

}

func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting all tasks"))
}
