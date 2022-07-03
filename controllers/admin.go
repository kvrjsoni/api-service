package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/helpers"
	"github.com/kvrjsoni/api-service/models"
)

// POST /admin/token/generate
func GenerateToken(c *gin.Context) {
	var createTokenInput models.CreateTokenInput
	// Validate request body
	if err := c.ShouldBindJSON(&createTokenInput); err != nil {
		// TODO - add error logging here
		helpers.DefaultApiResponseObject(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO - add error logging here
		fmt.Println("error while reading request body: ", err.Error())
	} else {
		// TODO - log the request payload
		fmt.Println("Request Payload for Generate Token: ", string(jsonData))
	}
	createTokenResult := models.CreateNewToken(createTokenInput)
	if errCreateToken := createTokenResult.Error; errCreateToken != nil {
		fmt.Println("error while generating token: ", errCreateToken.Error())
		helpers.DefaultApiResponseObject(c, http.StatusInternalServerError, gin.H{"error": errCreateToken.Error()})
	}
	helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": createTokenResult.Token})
}

// POST /admin/token/revoke
func RevokeToken(c *gin.Context) {
	var revokeTokenInput models.RevokeTokenInput
	// Validate request body
	if err := c.ShouldBindJSON(&revokeTokenInput); err != nil {
		// TODO - add error logging here
		helpers.DefaultApiResponseObject(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO - add error logging here
		fmt.Println("error while reading request body: ", err.Error())
	} else {
		// TODO - log the request payload
		fmt.Println("Request Payload for Revoke Token: ", string(jsonData))
	}

	revokeTokenResult := models.RevokeToken(revokeTokenInput.Token)
	if errRevokeToken := revokeTokenResult.Error; errRevokeToken != nil {
		fmt.Println("error while revoking token: ", errRevokeToken.Error())
		helpers.DefaultApiResponseObject(c, http.StatusInternalServerError, gin.H{"error": errRevokeToken.Error()})
	}
	helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": revokeTokenResult.Value})

}

// GET /admin/tokens
func ListAllTokens(c *gin.Context) {
	revokeTokenResult := models.GetAllTokens()
	if errRevokeToken := revokeTokenResult.Error; errRevokeToken != nil {
		fmt.Println("error while revoking token: ", errRevokeToken.Error())
		helpers.DefaultApiResponseObject(c, http.StatusInternalServerError, gin.H{"error": errRevokeToken.Error()})
	}
	helpers.DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": revokeTokenResult.Value})
}
