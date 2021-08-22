package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tonsV2/todo-go/pgk/di"
)

func GetEngine(environment di.Environment) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	baseUrl := environment.Configuration.BaseUrl

	router := r.Group(baseUrl)
	//router.GET("/health", environment.HealthHandler.GetHealthStatus)
	router.POST("/signup", environment.UserHandler.Signup)

	return r
}
