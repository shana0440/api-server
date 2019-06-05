package models

import (
	"github.com/u-job/api-server/db/entities"
)

// GetUserRequest use to get the request of GetUser handler
type GetUserRequest struct {
	ID int64 `route:"id"`
}

// UserResponse use to response user information
type UserResponse struct {
	User *entities.User `json:"user"`
}
