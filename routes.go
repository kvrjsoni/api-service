package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/helpers"
)

func initializeRoutes() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	router.POST("/token/generate", helpers.GenerateToken)
	router.Run(":3001")
}

func defaultRoute(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to API service")
}
