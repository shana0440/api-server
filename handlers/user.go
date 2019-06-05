package handlers

import (
	"net/http"

	"github.com/dannypsnl/rocket"
	"github.com/dannypsnl/rocket/response"
	log "github.com/sirupsen/logrus"

	"github.com/u-job/api-server/db/repositories"
	"github.com/u-job/api-server/models"
)

// CreateUser a example for how to use rocket
var CreateUser = rocket.Post("/users", func() *response.Response {
	user, err := repositories.CreateUser("u.job@demo.com")
	if err != nil {
		log.Error("Create user failed", err)
		return models.NewResponse().JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "Oops, something error",
		})
	}
	return models.NewResponse().JSON(http.StatusOK, models.UserResponse{
		User: user,
	})
})

// GetUser a example for how to use rocket
var GetUser = rocket.Get("/users/:id", func(request *models.GetUserRequest) *response.Response {
	user, err := repositories.GetUser(request.ID)
	if err != nil {
		log.Error("Get user failed", err)
		return models.NewResponse().JSON(http.StatusNotFound, models.ErrorResponse{
			Message: "cannot found the user",
		})
	}
	return models.NewResponse().JSON(http.StatusOK, models.UserResponse{
		User: user,
	})
})
