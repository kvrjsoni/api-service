package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUserName := "kvrjsoni"
	dbPassword := "nua@123"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "api_service"
	dsnString := (dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		// exit is db connection is not made
		panic("Failed to connect to database!")
	}

	autoMigrateError := db.AutoMigrate(
		&AdminUser{},
		&Token{},
	)

	if autoMigrateError != nil {
		panic("Failed to automigrate database tables")
	}

	DB = db
}
