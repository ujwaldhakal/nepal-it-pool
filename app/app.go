package main

import (
	"github.com/gin-gonic/gin"
	)

func main() {

	routes := gin.Default()
	registerDeveloperRoutes(routes)
	routes.Run() // run in 8080 port

}
