package main

import (
	"auth-service/internal/database"
	"auth-service/internal/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	database.InitDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	{
		userGroup := r.Group("/users")
		user.RegisterUserRoutes(userGroup)
	}

	err := r.Run()

	if err != nil {
		panic("it was failed to run the api")
	}
}
