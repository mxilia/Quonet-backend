package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	threadHandler "github.com/mxilia/Conflux-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Conflux-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Conflux-backend/internal/thread/usecase"

	userHandler "github.com/mxilia/Conflux-backend/internal/user/handler/rest"
	userRepository "github.com/mxilia/Conflux-backend/internal/user/repository"
	userUseCase "github.com/mxilia/Conflux-backend/internal/user/usecase"

	"github.com/mxilia/Conflux-backend/pkg/config"
)

func RegisterPublicRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {

	threadRepo := threadRepository.NewGormThreadRepository(db)
	threadUseCase := threadUseCase.NewThreadService(threadRepo)
	threadHandler := threadHandler.NewHttpThreadHandler(threadUseCase)

	userRepo := userRepository.NewGormUserRepository(db)
	userUseCase := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userUseCase, cfg)

	api := app.Group("/api/v2")

	authGroup := api.Group("/auth")

	googleAuthGroup := authGroup.Group("/google")

	googleAuthGroup.Get("/login", userHandler.GoogleLogin)
	googleAuthGroup.Get("/callback", userHandler.GoogleCallBack)

	threadGroup := api.Group("/threads")

	threadGroup.Post("/", threadHandler.CreateThread)
	threadGroup.Get("/", threadHandler.FindAllThreads)
	threadGroup.Get("/:id", threadHandler.FindThreadByID)
	threadGroup.Delete("/:id", threadHandler.DeleteThread)

	userGroup := api.Group("/users")

	userGroup.Get("/", userHandler.FindAllUsers)
	userGroup.Get("/:id", userHandler.FindUserByID)
	userGroup.Get("/handler/:handler", userHandler.FindUserByHandler)
	userGroup.Get("/email/:email", userHandler.FindUserByEmail)
	userGroup.Patch("/:id", userHandler.PatchUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)
}
