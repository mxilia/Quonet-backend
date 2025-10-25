package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	threadHandler "github.com/mxilia/Conflux-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Conflux-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Conflux-backend/internal/thread/usecase"
)

func RegisterPublicRoutes(app *fiber.App, db *gorm.DB) {

	threadRepo := threadRepository.NewGormThreadRepository(db)
	threadUseCase := threadUseCase.NewThreadService(threadRepo)
	threadHandler := threadHandler.NewHttpThreadHandler(threadUseCase)

	api := app.Group("/api")

	threadGroup := api.Group("/threads")

	threadGroup.Post("/", threadHandler.CreateThread)
	threadGroup.Get("/", threadHandler.GetAllThreads)
	threadGroup.Get("/:id", threadHandler.GetThreadByID)
	threadGroup.Delete("/:id", threadHandler.DeleteThread)
}
