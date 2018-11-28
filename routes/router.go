package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/imbarwinata/go-rest-core-v1/app/controllers"
	// "gitlab.com/imbarwinata/go-rest-core-v1/app/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
      userGroup.GET("/", user.Gets)
      userGroup.GET("/:id", user.Get)
      userGroup.POST("/", user.Insert)
      userGroup.PATCH("/:id/update", user.Update)
      userGroup.DELETE("/:id/delete", user.Delete)
		}
    peopleGroup := v1.Group("people")
		{
			user := new(controllers.UserController)
			peopleGroup.GET("/", user.Gets)
		}
	}
	return router

}
