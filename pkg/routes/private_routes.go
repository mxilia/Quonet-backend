package routes

import (
	"github.com/gofiber/fiber/v2"

	commentHandler "github.com/mxilia/Quonet-backend/internal/comment/handler/rest"
	commentRepository "github.com/mxilia/Quonet-backend/internal/comment/repository"
	commentUseCase "github.com/mxilia/Quonet-backend/internal/comment/usecase"
	"github.com/mxilia/Quonet-backend/internal/transaction"

	likeHandler "github.com/mxilia/Quonet-backend/internal/like/handler/rest"
	likeRepository "github.com/mxilia/Quonet-backend/internal/like/repository"
	likeUseCase "github.com/mxilia/Quonet-backend/internal/like/usecase"

	postHandler "github.com/mxilia/Quonet-backend/internal/post/handler/rest"
	postRepository "github.com/mxilia/Quonet-backend/internal/post/repository"
	postUseCase "github.com/mxilia/Quonet-backend/internal/post/usecase"

	sessionHandler "github.com/mxilia/Quonet-backend/internal/session/handler/rest"
	sessionRepository "github.com/mxilia/Quonet-backend/internal/session/repository"
	sessionUseCase "github.com/mxilia/Quonet-backend/internal/session/usecase"

	threadHandler "github.com/mxilia/Quonet-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Quonet-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Quonet-backend/internal/thread/usecase"

	userHandler "github.com/mxilia/Quonet-backend/internal/user/handler/rest"
	userRepository "github.com/mxilia/Quonet-backend/internal/user/repository"
	userUseCase "github.com/mxilia/Quonet-backend/internal/user/usecase"

	"github.com/mxilia/Quonet-backend/pkg/config"
	"github.com/mxilia/Quonet-backend/pkg/middleware"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {

	/* === Dependencies Wiring === */

	txManager := transaction.NewGormTxManager(db)

	threadRepo := threadRepository.NewGormThreadRepository(db)
	threadUseCase := threadUseCase.NewThreadService(threadRepo)
	threadHandler := threadHandler.NewHttpThreadHandler(threadUseCase)

	sessionRepo := sessionRepository.NewGormSessionRepository(db)
	sessionUseCase := sessionUseCase.NewSessionService(sessionRepo)

	userRepo := userRepository.NewGormUserRepository(db)
	userUseCase := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userUseCase, sessionUseCase, cfg)

	sessionHandler := sessionHandler.NewHttpSessionHandler(sessionUseCase, userUseCase, cfg)

	postRepo := postRepository.NewGormPostRepository(db)
	postUseCase := postUseCase.NewPostService(postRepo)
	postHandler := postHandler.NewHttpPostHandler(postUseCase)

	commentRepo := commentRepository.NewGormCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentService(commentRepo)
	commentHandler := commentHandler.NewHttpCommentHandler(commentUseCase)

	likeRepo := likeRepository.NewGormLikeRepository(db)
	likeUseCase := likeUseCase.NewLikeService(likeRepo, txManager, postUseCase, commentUseCase)
	likeHandler := likeHandler.NewHttpLikeHandler(likeUseCase)

	/* === Routes === */

	api := app.Group("/api/v2", middleware.JWTMiddleware(cfg, sessionUseCase, userUseCase))

	api.Get("/me", userHandler.GetUser)

	threadGroup := api.Group("/threads")

	threadGroup.Post("/", middleware.RequireAdmin(), threadHandler.CreateThread)
	threadGroup.Delete("/:id", middleware.RequireAdmin(), threadHandler.DeleteThread)

	sessionGroup := api.Group("/sessions")

	sessionGroup.Patch("/:email", middleware.RequireAdmin(), sessionHandler.RevokeToken)

	userGroup := api.Group("/users")

	userGroup.Patch("/:id", middleware.RequireUser(), userHandler.PatchUser)
	userGroup.Delete("/:id", middleware.RequireUser(), userHandler.DeleteUser)

	postGroup := api.Group("/posts")

	postGroup.Post("/", postHandler.CreatePost)
	postGroup.Get("/private", middleware.RequireUser(), postHandler.FindPrivatePosts)
	postGroup.Get("/private/:id", middleware.RequireUser(), postHandler.FindPrivatePostByID)
	postGroup.Patch("/:id", middleware.RequireUser(), postHandler.PatchPost)
	postGroup.Delete("/:id", middleware.RequireUser(), postHandler.DeletePost)

	likeGroup := api.Group("/likes")

	likeGroup.Post("/", likeHandler.CreateLike)
	likeGroup.Delete("/:id", middleware.RequireUser(), likeHandler.DeleteLike)

	commentGroup := api.Group("/comments")

	commentGroup.Post("/", commentHandler.CreateComment)
	commentGroup.Patch("/:id", middleware.RequireUser(), commentHandler.PatchComment)
	commentGroup.Delete("/:id", middleware.RequireUser(), commentHandler.DeleteComment)
}
