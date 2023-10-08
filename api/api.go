package api

import (
	"appname/api/activity"
	"appname/api/auth"
	"appname/api/task"
	"appname/api/template"
	"appname/api/user"
	"appname/conf"
	"appname/middleware"
	"fmt"
	"net/http"

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
	priv.POST("/activity", activity.Create)
	priv.PUT("/activity", activity.Update)
	priv.GET("/activity/:id", activity.Get)
	priv.POST("/activity/search", activity.Search)
	priv.GET("/self", auth.Self)
	priv.DELETE("/activity", activity.Delete)
	priv.POST("/task", task.Create)
	priv.PUT("/task", task.Update)
	priv.GET("/task/:id", task.Get)
	priv.POST("/task/search", task.Search)
	priv.POST("/task/complete", task.Complete)
	priv.DELETE("/task", task.Delete)
	priv.POST("/task/stats", task.Stats)
	priv.POST("/task/completedWeekPercentage", task.CompletedWeekPercentage)
	priv.POST("/user", user.Create)
	priv.POST("/user/:id/completeTutorial", user.CompleteTutorial)
	priv.GET("/user/:id", user.Get)
	priv.POST("/user/search", user.Search)
	priv.DELETE("/user/:id", user.Delete)
	priv.POST("/template", template.Create)
	priv.PUT("/template", template.Update)
	priv.GET("/template/:id", template.Get)
	priv.POST("/template/search", template.Search)
	priv.DELETE("/template", template.Delete)
	priv.POST("/template/:id/tasks", template.CreateTask)
	priv.DELETE("/template/tasks", template.DeleteTasks)
	priv.POST("/template/:id/apply", template.ApplyTemplate)
	return api
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Everything is FINE"})
}
