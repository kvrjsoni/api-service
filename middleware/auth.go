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

func AuthenticateAdminUser(c *gin.Context) {
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

func isPasswordValid(userNameFromAPI string, passwordFromAPI string) bool {
	loginData := models.GetLoginDetails(userNameFromAPI)
	err := bcrypt.CompareHashAndPassword([]byte(loginData.PasswordHash), []byte(passwordFromAPI))
	fmt.Println(err) // nil means it is a match
	return err == nil
}
