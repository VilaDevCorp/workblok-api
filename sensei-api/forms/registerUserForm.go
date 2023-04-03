package forms

type RegisterUserForm struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
