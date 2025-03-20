package app

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/null-bd/logger"
	"github.com/null-bd/staff-service-api/config"

	"github.com/null-bd/staff-service-api/internal/health"
	"github.com/null-bd/staff-service-api/internal/rest"
	"github.com/null-bd/staff-service-api/internal/staff"
	"github.com/null-bd/staff-service-api/internal/user"
)

type Application struct {
	HealthHandler rest.IHealthHandler
	StaffHandler  rest.IStaffHandler
	UserHandler   rest.IUserHandler
	DB            *pgxpool.Pool
	Config        *config.Config
}

func NewApplication(logger logger.Logger, cfg *config.Config, db *pgxpool.Pool) *Application {
	// Initialize repositories
	healthRepo := health.NewHealthRepository(db, logger)
	staffRepo := staff.NewStaffRepository(db, logger)
	userRepo := user.NewUserRepository(db, logger)

	// Initialize services
	healthSvc := health.NewHealthService(healthRepo, logger)
	staffSvc := staff.NewStaffService(staffRepo, logger)
	userSvc := user.NewUserService(userRepo, logger)

	// Initialize handler
	h := rest.NewHealthHandler(healthSvc, logger)
	s := rest.NewStaffHandler(staffSvc, logger)
	u := rest.NewUserHandler(userSvc, logger)

	return &Application{
		HealthHandler: h,
		StaffHandler:  s,
		UserHandler:   u,
		DB:            db,
		Config:        cfg,
	}
}
