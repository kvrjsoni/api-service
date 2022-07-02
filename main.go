package main

import (
	"github.com/kvrjsoni/api-service/models"
)

func main() {
	models.ConnectDatabase()

	// password := []byte("hahaha")
	// // Hashing the password with the default cost of 10
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(hashedPassword))

	// // Comparing the password with the hash
	// err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	// fmt.Println(err) // nil means it is a match

	initializeRoutes()
}
