package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/pkg/config"
	"github.com/mxilia/Conflux-backend/pkg/database"
	"github.com/mxilia/Conflux-backend/pkg/middleware"
	"github.com/mxilia/Conflux-backend/pkg/routes"
	"gorm.io/gorm"
)

func setupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.GetConfig(env)

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, nil, err
	}

	if env == "example" {
		db.Migrator().DropTable(&entities.Thread{})
	}
	if err := db.AutoMigrate(&entities.Thread{}); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}

func setupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app, cfg)
	routes.RegisterPublicRoutes(app, db)
	routes.RegisterPrivateRoutes(app)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}
