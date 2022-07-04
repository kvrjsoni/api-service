package models

import (
	"fmt"

	"github.com/kvrjsoni/api-service/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// make a connection to the db and also creates any tables if needed
func ConnectDatabase() bool {
	dbUserName := util.LoadEnvVariable("DB_USERNAME")
	dbPassword := util.LoadEnvVariable("DB_PASSWORD")
	dbHost := util.LoadEnvVariable("DB_HOST")
	dbPort := util.LoadEnvVariable("DB_PORT")
	dbName := util.LoadEnvVariable("DB_NAME")
	dsnString := (dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		// exit if db connection is not made
		fmt.Println("database connection failed")
		panic("Failed to connect to database!")
	}

	autoMigrateError := db.AutoMigrate(
		&AdminUser{},
		&Token{},
	)

	if autoMigrateError != nil {
		fmt.Println("db tables, auto migrate error")
		panic("Failed to automigrate database tables")
	}

	DB = db
	return true
}
