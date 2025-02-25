package rest

import (
	"github.com/null-bd/logger"
	"github.com/null-bd/microservice-name/internal/domain"
)

type IDomainHandler interface {
}

type domainHandler struct {
	domSvc domain.IDomainService
	log    logger.Logger
}

func NewDomainHandler(domSvc domain.IDomainService, logger logger.Logger) IDomainHandler {
	return &domainHandler{
		domSvc: domSvc,
		log:    logger,
	}
}
