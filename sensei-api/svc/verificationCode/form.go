package verificationCode

type CreateForm struct {
	Type string `json:"type" binding:"required"`
	Mail string `json:"mail" binding:"required"`
}

type UseForm struct {
	Type    string `json:"type" binding:"required"`
	Mail    string `json:"mail" binding:"required"`
	Code    string `json:"code" binding:"required"`
	NewPass string `json:"newPass"`
}
