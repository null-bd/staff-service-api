package rest

import (
	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/internal/user"
)

type IUserHandler interface {
}

type userHandler struct {
	userSvc user.IUserService
	log     logger.Logger
}

func NewUserHandler(userSvc user.IUserService, logger logger.Logger) IUserHandler {
	return &userHandler{
		userSvc: userSvc,
		log:     logger,
	}
}
