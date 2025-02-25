package user

import (
	"github.com/null-bd/logger"
)

type IUserService interface {
}

type userService struct {
	repo IUserRepository
	log  logger.Logger
}

func NewUserService(repo IUserRepository, logger logger.Logger) IUserService {
	return &userService{
		repo: repo,
		log:  logger,
	}
}
