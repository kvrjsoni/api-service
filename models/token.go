package models

import (
	"fmt"
	"time"

	"github.com/kvrjsoni/api-service/helpers"
)

type Token struct {
	ID         uint      `json:"id"`
	Token      string    `json:"token"`
	Status     bool      `json:"status"`
	ClientName string    `json:"client_name"`
	ExpireAt   time.Time `json:"expire_at"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at`
}

type CreateTokenInput struct {
	ClientName string `json:"client_name" binding:"required"`
}

type RevokeTokenInput struct {
	Token string `json:"token" binding:"required"`
}

type TokenReturnStruct struct {
	Value Token
	Error error
}

func CreateNewToken(createTokenInput CreateTokenInput) TokenReturnStruct {
	tokenDetails := Token{
		Token:      helpers.GenerateSecureToken(12),
		Status:     true,
		ClientName: createTokenInput.ClientName,
		ExpireAt:   helpers.AddDaysToCurrentTime(30),
	}

	fmt.Printf("%+v\n", tokenDetails)
	var dbCreateError error = nil

	if dbc := DB.Create(&tokenDetails); dbc.Error != nil {
		// DB create failed
		dbCreateError = dbc.Error
	}

	createTokenReturnData := TokenReturnStruct{
		Value: tokenDetails,
		Error: dbCreateError,
	}
	return createTokenReturnData
}

func RevokeToken(token string) TokenReturnStruct {
	var tokenDetails Token
	DB.Where("token = ?", token).Find(&tokenDetails)

	updateTokenDetails := map[string]interface{}{
		"updated_at": time.Now(),
		"status":     false,
	}
	var dbUpdateError error = nil
	query := DB.Debug().Model(&tokenDetails).Where("token = ?", token).Omit("id", "created_at", "deleted_at").Updates(updateTokenDetails)

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
