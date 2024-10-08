package auth

type LoginForm struct {
	Email      string `json:"email"      binding:"required"`
	Password   string `json:"password"   binding:"required"`
	RememberMe bool   `json:"rememberMe"`
	UserAgent  string `json:"userAgent"`
}

type SignUpForm struct {
	UserName string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LogoutForm struct {
	SessionId string `json:"sessionId" binding:"required"`
}
