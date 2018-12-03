package server

import (
	"github.com/gin-gonic/gin"
	"github.com/imbarwinata/go-gin-gorm-bolerplate/app/controllers"
	"github.com/imbarwinata/go-gin-gorm-bolerplate/app/middleware"
)

func NewRouter() *gin.Engine {
	// Register Controllers
	account := new(controllers.AccountController)
	article := new(controllers.ArticleController)
	auth := new(controllers.AuthController)
	health := new(controllers.HealthController)
	user := new(controllers.UserController)
	// Declare Middleware
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())

	router.GET("/health", health.Status)
	v1 := router.Group("v1")
	{
		authGroup := v1.Group("auth")
		{
			authGroup.POST("/", auth.Auth)
			authGroup.POST("/check-token", auth.AuthCheckToken)
			authGroup.POST("/register", user.Insert)
		}
		// Authentication required
		authorized := v1.Group("/")
		authorized.Use(middleware.JWTMiddleware())
		{
			authorized.GET("/articles", article.GetsAll)
			userGroup := authorized.Group("user")
			{
				userGroup.GET("/", user.Gets)
				userGroup.GET("/:id", user.Get)
				userGroup.PATCH("/:id/update", user.Update)
				userGroup.DELETE("/:id/delete", user.Delete)
				// User have articles
				articleGroup := userGroup.Group(":id/article")
				{
					articleGroup.GET("/", article.Gets)
					articleGroup.GET("/:articleid", article.Get)
					articleGroup.POST("/", article.Insert)
					articleGroup.PATCH("/:articleid/update", article.Update)
					articleGroup.DELETE("/:articleid/delete", article.Delete)
				}
				// User have account
				accountGroup := userGroup.Group(":id/account")
				{
					accountGroup.GET("/", account.Get)
					accountGroup.POST("/", account.Insert)
					accountGroup.PATCH("/update", account.Update)
				}
			}
			peopleGroup := authorized.Group("people")
			{
				peopleGroup.GET("/", user.Gets)
			}
		}
	}
	return router

}
