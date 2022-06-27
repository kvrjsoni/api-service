package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateToken(c *gin.Context) {
	fmt.Println("generating token..")
	token := GenerateSecureToken(12)
	DefaultApiResponseObject(c, http.StatusOK, gin.H{"data": token})
}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
