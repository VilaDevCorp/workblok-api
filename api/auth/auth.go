package auth

import (
	"fmt"
	"net/http"
	"workblok/conf"
	"workblok/mail"
	"workblok/svc"
	"workblok/svc/auth"
	"workblok/svc/verificationCode"
	"workblok/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	conf := conf.Get()
	var form auth.LoginForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	res, tokenString := svc.Auth.Login(c.Request.Context(), form)
	if res.Status == http.StatusOK {
		c.SetSameSite(http.SameSiteLaxMode)
		if conf.Env == "prod" {
			c.SetCookie("JWT_TOKEN", *tokenString, 30*24*60*60*1000, "/", conf.Prod.CookieHost, conf.Prod.CookieSecure, conf.Prod.CookieHttpOnly)
		} else {
			c.SetCookie("JWT_TOKEN", *tokenString, 30*24*60*60*1000, "/", conf.Dev.CookieHost, conf.Dev.CookieSecure, conf.Dev.CookieHttpOnly)
		}
	}
	c.JSON(res.Status, res.Result)
}

func SignUp(c *gin.Context) {
	var form auth.SignUpForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	res := svc.Auth.SignUp(c.Request.Context(), form)
	c.JSON(res.Status, res.Result)
}

func CreateVerificationCode(c *gin.Context) {
	conf := conf.Get()
	var form verificationCode.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	verificationCode, err := svc.VerificationCode.Create(c.Request.Context(), form)
	hostUrl := conf.Dev.FrontUrl
	if conf.Env == "prod" {
		hostUrl = conf.Prod.FrontUrl
	}
	if conf.Mail.Enabled {
		if verificationCode.Type == utils.VALIDATION_TYPE {
			err = mail.SendMail(form.Email, conf.Mail.MailTypes.AccountActivation.Subject,
				fmt.Sprintf(conf.Mail.MailTypes.AccountActivation.Body,
					hostUrl, form.Email, verificationCode.Code))
		} else if verificationCode.Type == utils.RECOVER_TYPE {
			err = mail.SendMail(form.Email, conf.Mail.MailTypes.PasswordReset.Subject,
				fmt.Sprintf(conf.Mail.MailTypes.PasswordReset.Body,
					hostUrl, form.Email, verificationCode.Code))
		}
	}

	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}

	res := utils.OkCreated(verificationCode)
	c.JSON(res.Status, nil)
}

func UseVerificationCode(c *gin.Context) {
	var form verificationCode.UseForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	res := svc.VerificationCode.UseCode(c.Request.Context(), form)
	c.JSON(res.Status, res.Result)
}

func Self(c *gin.Context) {
	jwtCookie, err := c.Cookie("JWT_TOKEN")
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	claims, err := utils.ValidateToken(jwtCookie)
	if err != nil {
		res := utils.InternalError(err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	user, err := svc.User.Get(c.Request.Context(), claims.Id)
	if err != nil {
		res := utils.NotFoundEntity(claims.Id.String())
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	res := utils.OkGet(user)
	c.JSON(res.Status, res.Result)
}
