package main

import (
	"fmt"

	"github.com/kvrjsoni/api-service/models"
)

func main() {
	// connecting to the database
	fmt.Println("connecting to the database...")
	isConnected := models.ConnectDatabase()
	if isConnected {
		fmt.Println("connected to the database")
	}
	// initializing API routes
	fmt.Println("initializing API routes")
	initializeRoutes()
}
