package auth

import (
	"net/http"
	"workblok/conf"
	"workblok/customerrors"
	"workblok/svc"
	"workblok/svc/auth"
	"workblok/svc/verificationCode"
	"workblok/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginResponse struct {
	AuthToken string    `json:"authToken"`
	SessionId uuid.UUID `json:"sessionId"`
}

func setRefreshCookie(c *gin.Context, refreshToken string, rememberMe bool) {
	conf := conf.Get()
	c.SetSameSite(http.SameSiteLaxMode)

	maxAge := 30 * 60 // 30 minutes
	if rememberMe {
		maxAge = 30 * 24 * 60 * 60 // 30 days
	}
	if conf.Env == "prod" {
		c.SetCookie("REFRESH_TOKEN", refreshToken,
			maxAge, "/api/public/refresh-token", conf.Prod.CookieHost,
			conf.Prod.CookieSecure, conf.Prod.CookieHttpOnly)
	} else {
		c.SetCookie("REFRESH_TOKEN", refreshToken, maxAge, "/api/public/refresh-token",
			conf.Dev.CookieHost, conf.Dev.CookieSecure, conf.Dev.CookieHttpOnly)
	}
}

func removeRefreshCookie(c *gin.Context) {
	conf := conf.Get()
	c.SetSameSite(http.SameSiteLaxMode)
	if conf.Env == "prod" {
		c.SetCookie("REFRESH_TOKEN", "", -1, "/api/public/refresh-token", conf.Prod.CookieHost,
			conf.Prod.CookieSecure, conf.Prod.CookieHttpOnly)
	} else {
		c.SetCookie("REFRESH_TOKEN", "", -1, "/api/public/refresh-token", conf.Dev.CookieHost,
			conf.Dev.CookieSecure, conf.Dev.CookieHttpOnly)
	}
}

func Login(c *gin.Context) {
	var form auth.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	svc := svc.Get()
	form.UserAgent = c.Request.UserAgent()
	loginResponse, err := svc.Auth.Login(c.Request.Context(), form)
	var res utils.HttpResponse
	if err != nil {
		switch err.(type) {
		case *customerrors.EmptyFormFieldsError:
			res = utils.ErrorResponse(http.StatusBadRequest,
				utils.GetStringPointer("No username or password present"), nil)
		case *customerrors.InvalidCredentialsError:
			res = utils.ErrorResponse(http.StatusUnauthorized,
				utils.GetStringPointer("Invalid credentials"),
				utils.GetStringPointer(customerrors.INVALID_CREDENTIALS))
		case *customerrors.NotValidatedAccountError:
			res = utils.ErrorResponse(http.StatusForbidden,
				utils.GetStringPointer("Not validated account"),
				utils.GetStringPointer(customerrors.NOT_VALIDATED_ACCOUNT))
		default:
			res = utils.InternalError(err)
		}
	} else {
		setRefreshCookie(c, loginResponse.RefreshJWTToken, form.RememberMe)
		res = utils.SuccessResponse(&LoginResponse{AuthToken: loginResponse.AuthJWTToken, SessionId: loginResponse.SessionId})
	}
	c.JSON(res.Status, res.Body)
}

func SignUp(c *gin.Context) {
	var form auth.SignUpForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest,
			utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	svc := svc.Get()
	user, err := svc.Auth.SignUp(c.Request.Context(), form)
	var res utils.HttpResponse
	if err != nil {
		switch err.(type) {
		case *customerrors.EmptyFormFieldsError:
			res = utils.ErrorResponse(http.StatusBadRequest,
				utils.GetStringPointer("No username, email or password present"), nil)
		case *customerrors.UsernameAlreadyInUseError:
			res = utils.ErrorResponse(http.StatusConflict,
				utils.GetStringPointer("Username already exist"),
				utils.GetStringPointer(customerrors.USERNAME_ALREADY_IN_USE))
		case *customerrors.EmailAlreadyInUseError:
			res = utils.ErrorResponse(http.StatusConflict,
				utils.GetStringPointer("Email already exist"),
				utils.GetStringPointer(customerrors.EMAIL_ALREADY_IN_USE))
		default:
			res = utils.InternalError(err)
		}
	} else {
		res = utils.SuccessResponse(user)
	}
	c.JSON(res.Status, res.Body)
}

func ResendValidationCode(c *gin.Context) {
	email, _ := c.Params.Get("email")
	svc := svc.Get()
	err := svc.VerificationCode.Create(
		c.Request.Context(),
		verificationCode.CreateForm{Email: email, Type: utils.VALIDATION_TYPE},
		nil,
	)

	var res utils.HttpResponse
	if err != nil {
		res = utils.InternalError(err)
	} else {
		res = utils.SuccessResponse(nil)
	}
	c.JSON(res.Status, res.Body)
}

func SendForgottenPasswordCode(c *gin.Context) {
	email, _ := c.Params.Get("email")
	svc := svc.Get()
	err := svc.VerificationCode.Create(
		c.Request.Context(),
		verificationCode.CreateForm{Email: email, Type: utils.RECOVER_TYPE},
		nil,
	)
	var res utils.HttpResponse
	if err != nil {
		res = utils.InternalError(err)

	} else {
		res = utils.SuccessResponse(nil)
	}
	c.JSON(res.Status, res.Body)
}

func ValidateAccount(c *gin.Context) {
	email, _ := c.Params.Get("email")
	code, _ := c.Params.Get("code")
	form := verificationCode.UseForm{Email: email, Code: code, Type: utils.VALIDATION_TYPE}
	svc := svc.Get()
	err := svc.VerificationCode.UseCode(c.Request.Context(), form)
	var res utils.HttpResponse
	if err != nil {
		switch err.(type) {
		case *customerrors.AlreadyUsedValidationCodeError:
			res = utils.ErrorResponse(http.StatusConflict, utils.GetStringPointer("Validation code already used"),
				utils.GetStringPointer(customerrors.ALREADY_USED_VALIDATION_CODE))
		case *customerrors.ExpiredValidationCodeError:
			res = utils.ErrorResponse(http.StatusGone, utils.GetStringPointer("Validation code expired"),
				utils.GetStringPointer(customerrors.EXPIRED_VALIDATION_CODE))
		case *customerrors.IncorrectValidationCodeError:
			res = utils.ErrorResponse(http.StatusUnauthorized, utils.GetStringPointer("Invalid validation code"),
				utils.GetStringPointer(customerrors.INCORRECT_VALIDATION_CODE))
		default:
			res = utils.InternalError(err)
		}
	} else {
		res = utils.SuccessResponse(nil)
	}
	c.JSON(res.Status, res.Body)
}

func ResetPassword(c *gin.Context) {
	email, _ := c.Params.Get("email")
	code, _ := c.Params.Get("code")
	body, err := c.GetRawData()
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	newPassStr := string(body)

	if newPassStr == "" {
		res := utils.ErrorResponse(
			http.StatusBadRequest,
			utils.GetStringPointer("No new password present"),
			nil,
		)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}

	form := verificationCode.UseForm{
		Email:   email,
		Code:    code,
		Type:    utils.RECOVER_TYPE,
		NewPass: newPassStr,
	}
	svc := svc.Get()
	err = svc.VerificationCode.UseCode(c.Request.Context(), form)
	var res utils.HttpResponse
	if err != nil {
		switch err.(type) {
		case *customerrors.AlreadyUsedValidationCodeError:
			res = utils.ErrorResponse(http.StatusConflict, utils.GetStringPointer("Validation code already used"),
				utils.GetStringPointer(customerrors.ALREADY_USED_VALIDATION_CODE))
		case *customerrors.ExpiredValidationCodeError:
			res = utils.ErrorResponse(http.StatusGone, utils.GetStringPointer("Validation code expired"),
				utils.GetStringPointer(customerrors.EXPIRED_VALIDATION_CODE))
		case *customerrors.IncorrectValidationCodeError:
			res = utils.ErrorResponse(http.StatusUnauthorized, utils.GetStringPointer("Invalid validation code"),
				utils.GetStringPointer(customerrors.INCORRECT_VALIDATION_CODE))
		default:
			res = utils.InternalError(err)
		}
	} else {
		res = utils.SuccessResponse(nil)
	}
	c.JSON(res.Status, res.Body)
}

func Self(c *gin.Context) {
	userId := c.Request.Context().Value(utils.UserIdKey).(uuid.UUID)
	svc := svc.Get()
	user, err := svc.User.Get(c.Request.Context(), userId)
	if err != nil {
		res := utils.NotFound("user", userId.String())
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	res := utils.SuccessResponse(user)
	c.JSON(res.Status, res.Body)
}

func RefreshToken(c *gin.Context) {

	refreshToken, err := c.Cookie("REFRESH_TOKEN")
	if err != nil {
		res := utils.ErrorResponse(
			http.StatusBadRequest,
			utils.GetStringPointer("Not Refresh token present"),
			utils.GetStringPointer(customerrors.NOT_REFRESH_TOKEN),
		)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}

	newAuthToken, newRefreshToken, newSessionId, err := svc.Get().Auth.RefreshToken(
		c.Request.Context(),
		refreshToken,
		c.Request.UserAgent(),
	)

	var res utils.HttpResponse
	if err != nil {
		switch err.(type) {
		case *customerrors.InvalidTokenError:
			res = utils.ErrorResponse(http.StatusUnauthorized, utils.GetStringPointer(err.Error()),
				utils.GetStringPointer(customerrors.INVALID_TOKEN))
			removeRefreshCookie(c)

		case *customerrors.UserAgentNotMatchError:
			res = utils.ErrorResponse(http.StatusUnauthorized, utils.GetStringPointer("User agent not match"),
				utils.GetStringPointer(customerrors.USER_AGENT_NOT_MATCH))
			removeRefreshCookie(c)

		case *customerrors.TokenAlreadyUsedError:
			res = utils.ErrorResponse(http.StatusUnauthorized, utils.GetStringPointer("Token already used"),
				utils.GetStringPointer(customerrors.TOKEN_ALREADY_USED))
			removeRefreshCookie(c)

		default:
			res = utils.InternalError(err)
		}
	} else {
		refreshTokenClaims, _ := utils.ValidateToken(refreshToken)
		setRefreshCookie(c, newRefreshToken, refreshTokenClaims.RememberMe)
		res = utils.SuccessResponse(&LoginResponse{AuthToken: newAuthToken, SessionId: newSessionId})
	}
	c.JSON(res.Status, res.Body)
}

func Logout(c *gin.Context) {
	var form auth.LogoutForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.ErrorResponse(http.StatusBadRequest, utils.GetStringPointer(err.Error()), nil)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	uuidSessionId, err := uuid.Parse(form.SessionId)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Body)
		return
	}
	svc.Get().Auth.Logout(c.Request.Context(), uuidSessionId)
	removeRefreshCookie(c)
}
