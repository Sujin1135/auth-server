package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunServer() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()

	if err != nil {
		panic("it was failed to run the server")
	}

	return r
}
