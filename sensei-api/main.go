package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"vilacorp.com/sensei/config"
	"vilacorp.com/sensei/controllers"
	"vilacorp.com/sensei/database"
	"vilacorp.com/sensei/utils"
)

func RegisterRoutes(router *mux.Router) {
	publicRouter := router.PathPrefix("/api/public").Subrouter()
	privateRouter := router.PathPrefix("/api/private").Subrouter()
	privateRouter.Use(authMiddleware)
	publicRouter.HandleFunc("/register", controllers.Register).Methods("POST")
	publicRouter.HandleFunc("/login", controllers.Login).Methods("POST")
	privateRouter.HandleFunc("/activity", controllers.CreateActivity).Methods("POST")
	privateRouter.HandleFunc("/activity", controllers.UpdateActivity).Methods("PUT")
	privateRouter.HandleFunc("/activity/{id}", controllers.GetActivity).Methods("GET")
	privateRouter.HandleFunc("/activity", controllers.SearchActivities).Methods("GET")
	privateRouter.HandleFunc("/activity/{id}", controllers.DeleteActivity).Methods("DELETE")
	privateRouter.HandleFunc("/activity", controllers.DeleteActivities).Methods("DELETE")
	privateRouter.HandleFunc("/task", controllers.CreateTask).Methods("POST")
	privateRouter.HandleFunc("/task/{id}", controllers.GetTask).Methods("GET")
	privateRouter.HandleFunc("/task", controllers.UpdateTask).Methods("PUT")
	privateRouter.HandleFunc("/task/{id}", controllers.DeleteTask).Methods("DELETE")
	privateRouter.HandleFunc("/user/{id}/tasks/{startDate}", controllers.GetUserTasks).Methods("GET")
	privateRouter.HandleFunc("/user/{id}/dans", controllers.UpdateUserDans).Methods("PATCH")
	privateRouter.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	privateRouter.HandleFunc("/self", controllers.Self).Methods("GET")

}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://www.coollocalhost.com:3000")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "http://www.coollocalhost.com:3000")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, credentials, Content-Length, Accept-Encoding, X-API-CSRF, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Set-Cookie")

			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			csrf := req.Header.Get("X-API-CSRF")
			if csrf == "" {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403- CSRF not provided"))
				return
			}
			jwt, err := req.Cookie("JWT_TOKEN")
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403- JWT not provided"))
				return
			}
			tokenClaims, jwtError := utils.ValidateToken(jwt.Value)
			if jwtError != nil {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403- JWT invalid or expired"))
				return
			}
			if !utils.CompareHash(csrf, tokenClaims.Csrf) {
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte("403- Invalid CSRF token"))
				return
			}
			next.ServeHTTP(w, req)
		})
}

func main() {
	config.LoadAppConfig()
	database.Connect(config.AppConfig.ConnectionString)
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
	log.Println("/api/task (PUT)")
	log.Println("/api/task/{id} (DELETE)")
	log.Println("/api/user/{id}/tasks/{startDate} (GET)")
	log.Println("/api/user/{id}/dans?add=true&value=5 (PATCH)")
	log.Println("/api/user/{id} (GET)")

	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))

	// var activity []entities.Activity
	// error := database.Instance.Model(&entities.Activity{}).Preload("User").Find(&activity).Error
	// if error != nil {
	// 	fmt.Print("ERROR")
	// } else {
	// 	fmt.Print(activity)
	// }
}
