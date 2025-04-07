package http

import (
	"github.com/ilyasa1211/ecommerce-scale-server/backend/services/user/internal/application"
)

type UserHandler struct {
	service application.UserService
}

func NewUserHandler(service application.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
