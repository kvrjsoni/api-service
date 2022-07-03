package main

import (
	"github.com/kvrjsoni/api-service/models"
)

func main() {
	// connecting to the database
	models.ConnectDatabase()
	// initializing API routes
	initializeRoutes()
}
