package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"vilacorp.com/sensei/controllers"
	"vilacorp.com/sensei/database"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/activity", controllers.CreateActivity).Methods("POST")
	router.HandleFunc("/api/activity", controllers.UpdateActivity).Methods("PUT")
	router.HandleFunc("/api/activity/{id}", controllers.GetActivity).Methods("GET")
	router.HandleFunc("/api/activity", controllers.SearchActivities).Methods("GET")
	router.HandleFunc("/api/activity/{id}", controllers.DeleteActivity).Methods("DELETE")
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/task/{id}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/user/{id}/tasks/{startDate}", controllers.GetUserTasks).Methods("GET")
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func main() {
	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterRoutes(router)
	enableCORS(router)

	// Start the server
	log.Println("/api/activity (POST)")
	log.Println("/api/activity (PUT)")
	log.Println("/api/activity/{id} (GET)")
	log.Println("/api/activity/{id} (DELETE)")
	log.Println("/api/activity (GET)")
	log.Println("/api/task (POST)")
	log.Println("/api/task/{id} (GET)")
	log.Println("/api/task/{id} (DELETE)")
	log.Println("/api/user/{id}/tasks/{startDate} (GET)")
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

	// var activity []entities.Activity
	// error := database.Instance.Model(&entities.Activity{}).Preload("User").Find(&activity).Error
	// if error != nil {
	// 	fmt.Print("ERROR")
	// } else {
	// 	fmt.Print(activity)
	// }
}
