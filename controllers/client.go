package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/helpers"
)

// POST /client/login
func ClientLogin(c *gin.Context) {
	// if this func runs that means the login is successful
	helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": "login_successful"})
}
