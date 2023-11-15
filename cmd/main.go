package main

import (
	"auth-service/internal/api"
	"auth-service/internal/database"
)

func main() {
	database.InitDB()
	api.Run()
}
