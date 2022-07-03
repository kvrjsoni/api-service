package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func AddDaysToCurrentTime(numberOfDaysToAdd int) time.Time {
	currentDate := time.Now()
	return currentDate.AddDate(0, 0, numberOfDaysToAdd)
}

func GenerateTokenHash(token string) string {
	tokenHash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(tokenHash)
}
