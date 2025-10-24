package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mxilia/Conflux-backend/pkg/routes"
)

func setUpRestServer() *fiber.App {
	app := fiber.New()
	routes.RegisterPublicRoutes(app)
	return app
}
