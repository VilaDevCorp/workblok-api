package verificationCode

type CreateForm struct {
	Type  string `json:"type" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UseForm struct {
	Type    string `json:"type" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Code    string `json:"code" binding:"required"`
	NewPass string `json:"newPass"`
}
