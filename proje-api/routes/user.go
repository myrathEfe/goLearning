package routes

import (
	"proje-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/user", controller.ListUsers)
	v1.POST("/user", controller.CreateUser)
}
