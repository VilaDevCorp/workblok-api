package auth

import (
	"fmt"
	"net/http"
	"sensei/mail"
	"sensei/svc"
	"sensei/svc/auth"
	"sensei/svc/verificationCode"
	"sensei/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
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
		c.SetCookie("JWT_TOKEN", *tokenString, 30*24*60*60*1000, "/", "trainwithsensei.com", true, true)
	}
	fmt.Println(res.Status)
	fmt.Println(res.Result)
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
	var form verificationCode.CreateForm
	err := c.ShouldBind(&form)
	if err != nil {
		res := utils.BadRequest(form, err)
		c.AbortWithStatusJSON(res.Status, res.Result)
		return
	}
	svc := svc.Get()
	verificationCode, err := svc.VerificationCode.Create(c.Request.Context(), form)
	if verificationCode.Type == utils.VALIDATION_TYPE {
		err = mail.SendMail("davidvilas@gmail.com", fmt.Sprintf("To validate the account you can use the code %s",
			verificationCode.Code))
	} else if verificationCode.Type == utils.RECOVER_TYPE {
		err = mail.SendMail("davidvilas@gmail.com", fmt.Sprintf("You can use the code %s for changing your password",
			verificationCode.Code))
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
