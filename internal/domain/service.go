package domain

import (
	"github.com/null-bd/logger"
)

type IDomainService interface {
}

type domainService struct {
	repo IDomainRepository
	log  logger.Logger
}

func NewDomainService(repo IDomainRepository, logger logger.Logger) IDomainService {
	return &domainService{
		repo: repo,
		log:  logger,
	}
}
