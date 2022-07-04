package models

import (
	"fmt"
	"time"
)

// creates the `admin_user` which will store all admin related data
type AdminUser struct {
	ID           uint
	UserName     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// this func returns the data from admin table for a sepified user
func GetLoginDetails(userName string) AdminUser {
	adminData := AdminUser{}
	DB.Where("user_name = ?", userName).Find(&adminData)
	fmt.Printf("%+v\n", adminData)
	return adminData
}
