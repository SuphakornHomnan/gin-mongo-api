package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.GET("/users", controllers.GetAllUser())
	router.GET("/user/:userId", controllers.GetUserById())
	router.POST("/user", controllers.CreateUser())
	router.PUT("/user/:userId", controllers.EditUserById())
	router.DELETE("user/:userId", controllers.DeleteUser())
}
