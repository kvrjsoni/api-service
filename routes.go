package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/controllers"
	"github.com/kvrjsoni/api-service/middleware"
	"github.com/kvrjsoni/api-service/util"
)

func initializeRoutes() {
	router := gin.Default()
	router.GET("/", defaultRoute)
	router.POST("/admin/token/generate", middleware.AuthenticateAdminUser, controllers.GenerateToken)
	router.POST("/admin/token/revoke", middleware.AuthenticateAdminUser, controllers.RevokeToken)
	router.GET("/admin/tokens", middleware.AuthenticateAdminUser, controllers.ListAllTokens)
	router.POST("/client/login", middleware.AuthenticateClientUser, controllers.ClientLogin)
	router.POST("/client/token/validate", controllers.ValidateToken)
	router.Run(":" + util.LoadEnvVariable("SERVER_PORT"))
}

func defaultRoute(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to API service")
}
