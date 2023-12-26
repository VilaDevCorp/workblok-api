package middleware

import (
	"workblok/utils"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrf := c.Request.Header.Get("X-API-CSRF")
		if csrf == "" {
			res := utils.Forbidden("CSRF not provided", nil)
			c.AbortWithStatusJSON(res.Status, res.Result)
			return
		}
		jwt, err := c.Cookie("JWT_TOKEN")
		if err != nil {
			res := utils.Forbidden("JWT not provided", err)
			c.AbortWithStatusJSON(res.Status, res.Result)
			return
		}
		tokenClaims, jwtError := utils.ValidateToken(jwt)
		if jwtError != nil {
			res := utils.Forbidden("Invalid or expired JWT ", err)
			c.AbortWithStatusJSON(res.Status, res.Result)
			return
		}
		if !utils.CompareHash(csrf, tokenClaims.Csrf) {
			res := utils.Forbidden("Invalid CSRF", err)
			c.AbortWithStatusJSON(res.Status, res.Result)
			return
		}
		c.Next()
	}
}
