package user

import (
	"auth-service/internal/api"
	"auth-service/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func convertUserToUserProfile(user *user) *userProfile {
	return &userProfile{
		Base: api.Base{
			ID:        user.Model.ID,
			CreatedAt: user.Model.CreatedAt,
		},
		userCommon: userCommon{
			Name: user.Name,
		},
	}
}

func createUser(c *gin.Context) {
	var body userCreateRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &user{Name: body.Name}
	database.DB.Create(user)

	c.JSON(http.StatusCreated, convertUserToUserProfile(user))
}

func RegisterUserRoutes(r *gin.RouterGroup) {
	err := database.DB.AutoMigrate(&user{})
	if err != nil {
		panic("occurred an error when user table migrating")
	}

	r.POST("", createUser)
}
