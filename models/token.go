package models

import (
	"fmt"
	"time"

	"github.com/kvrjsoni/api-service/helpers"
	"golang.org/x/crypto/bcrypt"
)

// creates the `tokens` which will store all client token related data
type Token struct {
	ID         uint      `json:"id"`
	Token      string    `json:"token"`
	IsActive   bool      `json:"is_active"`
	ClientName string    `json:"client_name"`
	ExpireAt   time.Time `json:"expire_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// used to accept input for generate token API
type CreateTokenInput struct {
	ClientName string `json:"client_name" binding:"required"`
}

// used to accept input for revoke token, validate token
type TokenInput struct {
	Token string `json:"token" binding:"required"`
}

type CreateTokenReturnStruct struct {
	Value Token
	Token string
	Error error
}

type TokenReturnStruct struct {
	Value Token
	Error error
}

type TokensReturnStruct struct {
	Value []Token
	Error error
}

// creates a new token for a client and stores the token hash in the db
func CreateNewToken(createTokenInput CreateTokenInput) CreateTokenReturnStruct {
	token := helpers.GenerateSecureToken(12)
	tokenHashed := helpers.GenerateTokenHash(token)
	tokenDetails := Token{
		Token:      tokenHashed,
		IsActive:   true,
		ClientName: createTokenInput.ClientName,
		ExpireAt:   helpers.AddDaysToCurrentTime(30),
	}

	var dbCreateError error = nil

	if dbc := DB.Create(&tokenDetails); dbc.Error != nil {
		// DB create failed
		dbCreateError = dbc.Error
	}

	createTokenReturnData := CreateTokenReturnStruct{
		Value: tokenDetails,
		Token: token,
		Error: dbCreateError,
	}
	return createTokenReturnData
}

// finds the token to be revoked and marks the token status as inactive
func RevokeToken(token string) TokenReturnStruct {
	var tokenDetails, tokenDetailsToBeRevoked Token
	tokenDetailsStruct := GetActiveTokenDetails()
	tokenDetailsFromDb := tokenDetailsStruct.Value
	tokenId := 0
	for _, tokenData := range tokenDetailsFromDb {
		// checking the token against all the active, non-expired tokens
		if err := bcrypt.CompareHashAndPassword([]byte(tokenData.Token), []byte(token)); err == nil {
			tokenId = int(tokenData.ID)
			tokenDetailsToBeRevoked = tokenData
			break
		}
	}

	if tokenId != 0 {
		updateTokenDetails := map[string]interface{}{
			"id":          tokenDetailsToBeRevoked.ID,
			"token":       tokenDetailsToBeRevoked.Token,
			"is_active":   false,
			"client_name": tokenDetailsToBeRevoked.ClientName,
			"expire_at":   time.Now(),
			"created_at":  tokenDetailsToBeRevoked.CreatedAt,
			"updated_at":  time.Now(),
		}

		var dbUpdateError error = nil
		query := DB.Debug().Model(&tokenDetails).Where("id = ?", tokenId).Updates(updateTokenDetails)

		if dbError := query.Error; dbError != nil {
			fmt.Println("Error while revoking the template", dbError.Error())
			dbUpdateError = dbError
		}

		createTokenReturnData := TokenReturnStruct{
			Value: tokenDetails,
			Error: dbUpdateError,
		}
		return createTokenReturnData
	}
	return TokenReturnStruct{}
}

// returns all the token data from the db
func GetAllTokens() TokensReturnStruct {
	var tokenDetails []Token
	var listTokenError error = nil

	if listTokens := DB.Debug().Find(&tokenDetails); listTokens.Error != nil {
		listTokenError = listTokens.Error
	}

	return TokensReturnStruct{
		Value: tokenDetails,
		Error: listTokenError,
	}
}

// returns all the active client tokens
func GetActiveTokenDetails() TokensReturnStruct {
	var tokenDetails []Token
	var listTokenError error = nil

	if listTokens := DB.Debug().Where("is_active = ? ", true).Where("expire_at > ? ", time.Now()).Find(&tokenDetails); listTokens.Error != nil {
		listTokenError = listTokens.Error
	}

	return TokensReturnStruct{
		Value: tokenDetails,
		Error: listTokenError,
	}
}

// checks if a client token is valid or not
func IsTokenValid(token string) bool {
	tokenDetailsStruct := GetActiveTokenDetails()
	tokenDetails := tokenDetailsStruct.Value
	for _, tokenData := range tokenDetails {
		// checking the token against all the active, non-expired tokens
		if err := bcrypt.CompareHashAndPassword([]byte(tokenData.Token), []byte(token)); err == nil {
			return true
		}
	}
	return false
}
