package main

import (
	"proje-api/config"
	"proje-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":4545")
}
