package main

import (
	"github.com/gin-gonic/gin"
	)

func regRoutes(routes *gin.Engine) {

	routes.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})
}

func main() {

	routes := gin.Default()
	registerDeveloperRoutes(routes)
	routes.Run() // run in 8080 port

}
