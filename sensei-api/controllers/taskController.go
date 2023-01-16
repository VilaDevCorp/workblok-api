package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"vilacorp.com/sensei/database"
	"vilacorp.com/sensei/entities"
	"vilacorp.com/sensei/forms"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var taskForm forms.CreateTaskForm
	json.NewDecoder(r.Body).Decode(&taskForm)
	parsedDueDate, _ := time.Parse("2-1-2006", taskForm.DueDate)
	newTask := &entities.Task{DueDate: parsedDueDate, ActivityId: taskForm.ActivityId, UserId: taskForm.UserId}
	database.Instance.Create(&newTask)

	json.NewEncoder(w).Encode(newTask)
}

func checkIfTaskExists(taskId uuid.UUID) bool {
	var task entities.Task
	database.Instance.First(&task, taskId)
	nullUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if task.ID == nullUUID {
		return false
	}
	return true
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	taskIdStr := mux.Vars(r)["id"]
	taskId, _ := uuid.FromString(taskIdStr)
	if checkIfTaskExists(taskId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task Not Found!")
		return
	}
	var task entities.Task
	database.Instance.Model(&entities.Task{}).Preload("Activity").First(&task, taskId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetUserTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []entities.Task
	userId := mux.Vars(r)["id"]
	weekStart := mux.Vars(r)["startDate"]
	parsedWeekStart, _ := time.Parse("2-1-2006", weekStart)
	log.Println(fmt.Sprintf("%s", parsedWeekStart))
	// parsedWeekEnd := parsedWeekStart.Add(time.Hour * 24 * 7)
	database.Instance.Model(&entities.Task{}).Where("user_id = ? AND due_date > ?", userId, parsedWeekStart).Preload("Activity").Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	taskIdStr := mux.Vars(r)["id"]
	taskId, _ := uuid.FromString(taskIdStr)
	if checkIfTaskExists(taskId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task Not Found!")
		return
	}
	var task entities.Task
	database.Instance.Delete(&task, taskId)
	json.NewEncoder(w).Encode("Task Deleted Successfully!")
}
