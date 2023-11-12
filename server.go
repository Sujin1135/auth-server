package main

import (
	"auth-service/database"
	"auth-service/routers"
)

func main() {
	database.InitDB()
	_ = routers.RunServer()
}
