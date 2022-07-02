package models

import (
	"fmt"
	"time"
)

type AdminUser struct {
	ID           uint
	UserName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func GetLoginDetails(userName string) AdminUser {
	adminData := AdminUser{}
	DB.Where("user_name = ?", userName).Find(&adminData)
	fmt.Printf("%+v\n", adminData)
	return adminData
}
