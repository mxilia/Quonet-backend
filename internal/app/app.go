package app

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/pkg/config"
	"github.com/mxilia/Conflux-backend/pkg/database"
	"github.com/mxilia/Conflux-backend/pkg/middleware"
	"github.com/mxilia/Conflux-backend/pkg/routes"
	"gorm.io/gorm"
)

func setupOwner(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entities.User{}).Where("email = ?", os.Getenv("OWNER_EMAIL")).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	owner := &entities.User{Email: os.Getenv("OWNER_EMAIL"), Handler: os.Getenv("OWNER_HANDLER"), Role: "owner"}
	if err := db.Create(owner).Error; err != nil {
		return err
	}
	return nil
}

func setupDependencies(env string) (*gorm.DB, *config.Config, error) {
	cfg := config.GetConfig(env)

	db, err := database.Connect(cfg)
	if err != nil {
		return nil, nil, err
	}

	if env == "dev" {
		db.Migrator().DropTable(&entities.Thread{}, &entities.User{}, &entities.Post{}, &entities.Like{}, &entities.Comment{})
	}

	if err := db.AutoMigrate(&entities.Thread{}, &entities.User{}, &entities.Post{}, &entities.Like{}, &entities.Comment{}); err != nil {
		return nil, nil, err
	}

	if err := setupOwner(db); err != nil {
		return nil, nil, err
	}

	return db, cfg, nil
}

func setupRestServer(db *gorm.DB, cfg *config.Config) (*fiber.App, error) {
	app := fiber.New()
	middleware.FiberMiddleware(app, cfg)
	routes.RegisterPublicRoutes(app, db, cfg)
	routes.RegisterPrivateRoutes(app, db, cfg)
	routes.RegisterNotFoundRoute(app)
	return app, nil
}
