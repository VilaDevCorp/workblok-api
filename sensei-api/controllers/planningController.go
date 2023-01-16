package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"vilacorp.com/sensei/database"
	"vilacorp.com/sensei/entities"
)

func CreatePlanning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var planning entities.Planning
	json.NewDecoder(r.Body).Decode(&planning)
	database.Instance.Create(&planning)

	json.NewEncoder(w).Encode(planning)
}

func checkIfPlanningExists(planningId uuid.UUID) bool {
	var planning entities.Planning
	database.Instance.First(&planning, planningId)
	nullUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if planning.ID == nullUUID {
		return false
	}
	return true
}

func GetPlanning(w http.ResponseWriter, r *http.Request) {
	planningIdStr := mux.Vars(r)["id"]
	planningId, _ := uuid.FromString(planningIdStr)
	if checkIfPlanningExists(planningId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Planning Not Found!")
		return
	}
	var planning entities.Planning
	database.Instance.Model(&entities.Planning{}).Preload("Activity").First(&planning, planningId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(planning)
}

func DeletePlanning(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	planningIdStr := mux.Vars(r)["id"]
	planningId, _ := uuid.FromString(planningIdStr)
	if checkIfPlanningExists(planningId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Planning Not Found!")
		return
	}
	var planning entities.Planning
	database.Instance.Delete(&planning, planningId)
	json.NewEncoder(w).Encode("Planning Deleted Successfully!")
}
