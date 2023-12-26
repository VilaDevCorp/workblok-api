package api

import (
	"fmt"
	"net/http"
	"workblok/api/auth"
	"workblok/api/block"
	"workblok/api/user"
	"workblok/conf"
	"workblok/middleware"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Start() error {
	conf := conf.Get()
	r := GetRouter()
	log.Info().Msg(fmt.Sprintf("%s", conf.Port))
	err := r.Run(fmt.Sprintf("%s:%s", conf.IP, conf.Port))
	return err
}

func GetRouter() *gin.Engine {
	api := gin.Default()
	api.Use(gin.Recovery())
	api.Use(ginzerolog.Logger("gin"))
	api.Use(middleware.Cors())
	pub := api.Group("/api/public")
	pub.POST("/login", auth.Login)
	pub.POST("/useVerificationCode", auth.UseVerificationCode)
	pub.POST("/newVerificationCode", auth.CreateVerificationCode)
	pub.POST("/register", auth.SignUp)
	priv := api.Group("/api/private")
	priv.Use(middleware.Authentication())
	pub.GET("/health", health)
	priv.POST("/block", block.Create)
	priv.POST("/block/penalty", block.ApplyPenalty)
	priv.POST("/block/:id/finish", block.Finish)
	priv.GET("/block/:id", block.Get)
	priv.GET("/block/active", block.GetActive)
	priv.POST("/block/search", block.Search)
	priv.POST("/block/stats", block.Stats)
	priv.GET("/self", auth.Self)
	priv.DELETE("/block", block.Delete)
	priv.POST("/user", user.Create)
	priv.POST("/user/:id/completeTutorial", user.CompleteTutorial)
	priv.GET("/user/:id", user.Get)
	priv.POST("/user/search", user.Search)
	priv.DELETE("/user/:id", user.Delete)
	priv.PATCH("/user/conf", user.ConfUpdate)
	return api
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Everything is FINE"})
}
