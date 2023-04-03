package main

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
