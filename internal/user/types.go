package user

import (
	"auth-service/internal/api"
)

type userCommon struct {
	Name string `json:"name" binding:"required"`
}

type userProfile struct {
	api.Base
	userCommon
}

type userCreateRequest struct {
	userCommon
}
