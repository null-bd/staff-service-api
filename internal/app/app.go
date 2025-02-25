package app

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/logger"
	"github.com/null-bd/microservice-name/config"

	"github.com/null-bd/microservice-name/internal/domain"
	"github.com/null-bd/microservice-name/internal/health"
	"github.com/null-bd/microservice-name/internal/rest"
)

type Application struct {
	HealthHandler rest.IHealthHandler
	DomHandler    rest.IDomainHandler
	DB            *pgxpool.Pool
	Config        *config.Config
}

func NewApplication(logger logger.Logger, cfg *config.Config, db *pgxpool.Pool) *Application {
	// Initialize repositories
	healthRepo := health.NewHealthRepository(db, logger)
	domRepo := domain.NewDomainRepository(db, logger)

	// Initialize services
	healthSvc := health.NewHealthService(healthRepo, logger)
	domSvc := domain.NewDomainService(domRepo, logger)

	// Initialize handler
	h := rest.NewHealthHandler(healthSvc, logger)
	d := rest.NewDomainHandler(domSvc, logger)

	return &Application{
		HealthHandler: h,
		DomHandler:    d,
		DB:            db,
		Config:        cfg,
	}
}
