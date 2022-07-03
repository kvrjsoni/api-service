package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/helpers"
	"github.com/kvrjsoni/api-service/models"
)

// POST /client/login
func ClientLogin(c *gin.Context) {
	// if this func runs that means the login is successful
	helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": "login_successful"})
}

// POST /client/token/validate
func ValidateToken(c *gin.Context) {
	var validateTokenInput models.TokenInput
	// Validate request body
	if err := c.ShouldBindJSON(&validateTokenInput); err != nil {
		// TODO - add error logging here
		helpers.DefaultApiResponseObject(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isTokenValid := models.IsTokenValid(validateTokenInput.Token)
	if !isTokenValid {
		helpers.DefaultApiResponseObject(c, http.StatusUnauthorized, gin.H{"error": "invalid_token"})
	} else {
		helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": "token_valid"})
	}
}
