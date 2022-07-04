package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// generates a secure random token string with the specified length
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

// func is used to add `x` number for days to current time, used to calculate expiry date
func AddDaysToCurrentTime(numberOfDaysToAdd int) time.Time {
	currentDate := time.Now()
	return currentDate.AddDate(0, 0, numberOfDaysToAdd)
}

// generates a token hash for a given token string
func GenerateTokenHash(token string) string {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(tokenHash)
}
