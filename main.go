package main

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
