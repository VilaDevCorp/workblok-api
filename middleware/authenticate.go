package middleware

import (
	"context"
	"net/http"
	"strings"
	"workblok/customerrors"
	"workblok/utils"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		headerSplit := strings.Split(authHeader, " ")
		if len(headerSplit) != 2 {
			res := utils.ErrorResponse(
				http.StatusUnauthorized,
				utils.GetStringPointer("Not Auth token present"),
				utils.GetStringPointer(customerrors.NOT_AUTH_JWT_TOKEN),
			)
			c.AbortWithStatusJSON(res.Status, res.Body)
			return
		}
		if headerSplit[0] != "Bearer" {
			res := utils.ErrorResponse(
				http.StatusUnauthorized,
				utils.GetStringPointer("Not Auth token present"),
				utils.GetStringPointer(customerrors.NOT_AUTH_JWT_TOKEN),
			)
			c.AbortWithStatusJSON(res.Status, res.Body)
			return
		}
		authToken := headerSplit[1]
		if authToken == "" {
			res := utils.ErrorResponse(
				http.StatusUnauthorized,
				utils.GetStringPointer("Not Auth token present"),
				utils.GetStringPointer(customerrors.NOT_AUTH_JWT_TOKEN),
			)
			c.AbortWithStatusJSON(res.Status, res.Body)
			return
		}

		tokenClaims, jwtError := utils.ValidateToken(authToken)
		if jwtError != nil {
			res := utils.ErrorResponse(http.StatusUnauthorized,
				utils.GetStringPointer("Invalid or expired JWT"),
				utils.GetStringPointer(customerrors.INVALID_TOKEN),
			)
			c.AbortWithStatusJSON(res.Status, res.Body)
			return
		}
		newCtx := c.Request.Context()
		newCtx = context.WithValue(newCtx, utils.UserIdKey, tokenClaims.UserId)
		c.Request = c.Request.WithContext(newCtx)

		c.Next()
	}
}
