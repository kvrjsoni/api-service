package main

import "github.com/kvrjsoni/api-service/models"

func main() {
	models.ConnectDatabase()
	initializeRoutes()
}
