package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"vilacorp.com/sensei/database"
	"vilacorp.com/sensei/entities"
	"vilacorp.com/sensei/forms"
	"vilacorp.com/sensei/utils"
)

func hashPassword(password string) (result []byte, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func validateMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var registerForm forms.RegisterUserForm
	json.NewDecoder(r.Body).Decode(&registerForm)
	if registerForm.Username == "" || registerForm.Password == "" || registerForm.Mail == "" || !validateMail(registerForm.Mail) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400- Malformed request form"))
		return
	}
	var hashPassword, hashErr = hashPassword(registerForm.Password)
	if hashErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500- A server error has occurred"))
		return
	}
	newUser := &entities.User{Username: registerForm.Username, Mail: registerForm.Mail, Password: string(hashPassword[:])}
	if err := database.Instance.Create(&newUser).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500- A server error has occurred"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

type LoginResult struct {
	Csrf string `json:"csrf"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginForm forms.LoginForm
	json.NewDecoder(r.Body).Decode(&loginForm)
	if loginForm.Username == "" || loginForm.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400- Malformed request form"))
		return
	}
	var loginUser entities.User

	if err := database.Instance.Model(&entities.User{}).Where("username = ?", loginForm.Username).Find(&loginUser).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404- The user was not found"))
		return
	}

	if !checkPassword(loginUser.Password, loginForm.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401- The password is incorrect"))
		return
	}
	csrfToken, err := utils.GenerateRandomToken(64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500- A server error has occurred"))
		return
	}
	// hash csrf
	hashedCsrfToken, err := utils.HashAndSalt(csrfToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500- A server error has occurred"))
		return
	}

	tokenString, err := utils.GenerateJWT(loginUser.ID.String(), loginUser.Mail, loginUser.Username, csrfToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500- Internal error at login"))
		return
	}
	cookie := http.Cookie{}
	cookie.Name = "JWT_TOKEN"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(30 * 24 * time.Hour)
	cookie.Secure = false
	cookie.HttpOnly = false
	cookie.Domain = "coollocalhost.com:3000"
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Path = "/"
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusOK)
	result := LoginResult{Csrf: hashedCsrfToken}
	json.NewEncoder(w).Encode(result)
	return
}

func Self(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jwt, _ := r.Cookie("JWT_TOKEN")
	tokenClaims, _ := utils.ValidateToken(jwt.Value)
	userId := tokenClaims.Id
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user entities.User
	database.Instance.Model(&entities.User{}).First(&user, userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	return
}

func UpdateUserDans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userIdStr := mux.Vars(r)["id"]
	userId, _ := uuid.FromString(userIdStr)
	isAdd, _ := strconv.ParseBool(r.URL.Query().Get("add"))
	value, _ := strconv.Atoi(r.URL.Query().Get("value"))

	var oldUser entities.User

	database.Instance.Model(&entities.User{}).First(&oldUser, userId)
	var quantity = oldUser.Dans

	if isAdd {
		quantity = oldUser.Dans + value
	} else {
		quantity = oldUser.Dans - value
	}
	database.Instance.Model(oldUser).Update("Dans", quantity)

	json.NewEncoder(w).Encode(oldUser)
}

func UpdateUserDansInternal(userId uuid.UUID, isAdd bool, value int) {

	var oldUser entities.User

	database.Instance.Model(&entities.User{}).First(&oldUser, userId)
	var quantity = oldUser.Dans

	if isAdd {
		quantity = oldUser.Dans + value
		log.Println(fmt.Sprintf("Vamos a añadir %d dans", value))
	} else {
		quantity = oldUser.Dans - value
		log.Println(fmt.Sprintf("Vamos a quitar %d dans", value))
	}

	database.Instance.Model(oldUser).Update("Dans", quantity)
}

func checkIfUserExists(userId uuid.UUID) bool {
	var user entities.User
	database.Instance.First(&user, userId)
	nullUUID, _ := uuid.FromString("00000000-0000-0000-0000-000000000000")
	if user.ID == nullUUID {
		return false
	}
	return true
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := mux.Vars(r)["id"]
	userId, _ := uuid.FromString(userIdStr)
	if checkIfUserExists(userId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User Not Found!")
		return
	}
	var user entities.User
	database.Instance.Model(&entities.User{}).First(&user, userId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// func GetUserTasks(w http.ResponseWriter, r *http.Request) {
// 	var tasks []entities.Task
// 	userId := mux.Vars(r)["id"]
// 	weekStart := mux.Vars(r)["startDate"]
// 	parsedWeekStart, _ := time.Parse("2-1-2006", weekStart)
// 	log.Println(fmt.Sprintf("%s", parsedWeekStart))
// 	// parsedWeekEnd := parsedWeekStart.Add(time.Hour * 24 * 7)
// 	database.Instance.Model(&entities.Task{}).Where("user_id = ? AND due_date > ?", userId, parsedWeekStart).Preload("Activity").Find(&tasks)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	taskIdStr := mux.Vars(r)["id"]
// 	taskId, _ := uuid.FromString(taskIdStr)
// 	if checkIfTaskExists(taskId) == false {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode("Task Not Found!")
// 		return
// 	}
// 	var task entities.Task
// 	database.Instance.Delete(&task, taskId)
// 	json.NewEncoder(w).Encode("Task Deleted Successfully!")
// }
