package auth

import (
	"net/http"
	"sensei/svc"
	"sensei/svc/auth"
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
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("JWT_TOKEN", *tokenString, 30*24*60*60*1000, "/", "www.coollocalhost.com", false, false)
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
