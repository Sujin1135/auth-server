package user

import (
	"auth-service/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name string
}

func createUser(c *gin.Context) {
	user := &User{Name: "D42"}
	database.DB.Create(user)

	c.JSON(http.StatusCreated, user)
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic("occurred an error when user table migrating")
	}

	r.POST("", createUser)
}
