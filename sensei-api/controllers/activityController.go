package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"vilacorp.com/sensei/database"
	"vilacorp.com/sensei/entities"
)

func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activity, oldActivity entities.Activity
	json.NewDecoder(r.Body).Decode(&activity)
	database.Instance.Model(&entities.Activity{}).First(&oldActivity, activity.ID)

	database.Instance.Model(oldActivity).Updates(entities.Activity{Name: activity.Name, Size: activity.Size, Icon: activity.Icon})

	json.NewEncoder(w).Encode(activity)
}

func CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activity entities.Activity
	json.NewDecoder(r.Body).Decode(&activity)
	database.Instance.Create(&activity)

	json.NewEncoder(w).Encode(activity)
}

func checkIfActivityExists(activityId uuid.UUID) bool {
	var activity entities.Activity
	database.Instance.First(&activity, activityId)
	nullUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if activity.ID == nullUUID {
		return false
	}
	return true
}

func GetActivity(w http.ResponseWriter, r *http.Request) {
	activityIdStr := mux.Vars(r)["id"]
	activityId, _ := uuid.FromString(activityIdStr)
	if checkIfActivityExists(activityId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Activity Not Found!")
		return
	}
	var activity entities.Activity
	database.Instance.Model(&entities.Activity{}).Preload("User").First(&activity, activityId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activity)
}

func SearchActivities(w http.ResponseWriter, r *http.Request) {
	var activities []entities.Activity
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	search := r.URL.Query().Get("search")
	pagination := database.Pagination{Page: page}

	if search == "" {
		database.Instance.Scopes(database.Paginate(activities, &pagination, database.Instance)).Model(&entities.Activity{}).Order("creation_date").Preload("User").Find(&activities)
	} else {
		database.Instance.Scopes(database.Paginate(activities, &pagination, database.Instance)).Model(&entities.Activity{}).Preload("User").Where("name like ?", "%"+search+"%").Find(&activities)
	}

	pagination.Content = activities

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pagination)
}

func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	activityIdStr := mux.Vars(r)["id"]
	activityId, _ := uuid.FromString(activityIdStr)
	if checkIfActivityExists(activityId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Activity Not Found!")
		return
	}
	var activity entities.Activity
	database.Instance.Delete(&activity, activityId)
	json.NewEncoder(w).Encode("Activity Deleted Successfully!")
}
