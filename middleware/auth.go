package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kvrjsoni/api-service/helpers"
	"github.com/kvrjsoni/api-service/models"
	"golang.org/x/crypto/bcrypt"
)

// this function validates the admin user by checking the basic auth token provided with the API call
func AuthenticateAdminUser(c *gin.Context) {
	// TODO - add index check to prevent index out of range error
	basicAuthHeaderValue := c.Request.Header["Authorization"][0]
	basicAuthTrimmed := strings.Replace(basicAuthHeaderValue, "Basic ", "", -1)
	originalDetails, err := base64.StdEncoding.DecodeString(basicAuthTrimmed)
	if err != nil {
		fmt.Printf("%q\n", err)
	}
	userNameAndPasswordFromAPI := string(originalDetails)
	loginDetails := strings.Split(userNameAndPasswordFromAPI, ":")
	userName := loginDetails[0]
	password := loginDetails[1]
	if !isPasswordValid(userName, password) {
		// TODO - add logs here to capture unauthorized accesses
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		helpers.DefaultApiResponseObject(c, 401, gin.H{"error": "unauthorized access detected"})
		return
	}
}

// this func makes the check for the username and password received from the API call
func isPasswordValid(userNameFromAPI string, passwordFromAPI string) bool {
	loginData := models.GetLoginDetails(userNameFromAPI)
	err := bcrypt.CompareHashAndPassword([]byte(loginData.PasswordHash), []byte(passwordFromAPI))
	fmt.Println(err) // nil means it is a match
	return err == nil
}

// this function validates the client user by checking the api-key token provided with the API call
func AuthenticateClientUser(c *gin.Context) {
	// TODO - add index check to prevent index out of range error
	xApiKey := c.Request.Header["X-Api-Key"][0]
	isTokenValid := models.IsTokenValid(xApiKey)
	if !isTokenValid {
		// TODO - add logs here to capture unauthorized accesses
		c.Abort()
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		helpers.DefaultApiResponseObject(c, 401, gin.H{"error": "login_failed"})
		return
	}
}
