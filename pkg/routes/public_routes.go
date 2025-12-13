package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	threadHandler "github.com/mxilia/Quonet-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Quonet-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Quonet-backend/internal/thread/usecase"
	"github.com/mxilia/Quonet-backend/internal/transaction"

	userHandler "github.com/mxilia/Quonet-backend/internal/user/handler/rest"
	userRepository "github.com/mxilia/Quonet-backend/internal/user/repository"
	userUseCase "github.com/mxilia/Quonet-backend/internal/user/usecase"

	postHandler "github.com/mxilia/Quonet-backend/internal/post/handler/rest"
	postRepository "github.com/mxilia/Quonet-backend/internal/post/repository"
	postUseCase "github.com/mxilia/Quonet-backend/internal/post/usecase"

	commentHandler "github.com/mxilia/Quonet-backend/internal/comment/handler/rest"
	commentRepository "github.com/mxilia/Quonet-backend/internal/comment/repository"
	commentUseCase "github.com/mxilia/Quonet-backend/internal/comment/usecase"

	sessionHandler "github.com/mxilia/Quonet-backend/internal/session/handler/rest"
	sessionRepository "github.com/mxilia/Quonet-backend/internal/session/repository"
	sessionUseCase "github.com/mxilia/Quonet-backend/internal/session/usecase"

	likeHandler "github.com/mxilia/Quonet-backend/internal/like/handler/rest"
	likeRepository "github.com/mxilia/Quonet-backend/internal/like/repository"
	likeUseCase "github.com/mxilia/Quonet-backend/internal/like/usecase"

	"github.com/mxilia/Quonet-backend/pkg/config"
)

func RegisterPublicRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {

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

	api := app.Group("/api/v2")

	authGroup := api.Group("/auth")

	authGroup.Post("/logout", sessionHandler.Logout)

	googleAuthGroup := authGroup.Group("/google")

	googleAuthGroup.Get("/login", userHandler.GoogleLogin)
	googleAuthGroup.Get("/callback", userHandler.GoogleCallBack)

	threadGroup := api.Group("/threads")

	threadGroup.Get("/", threadHandler.FindAllThreads)
	threadGroup.Get("/:id", threadHandler.FindThreadByID)

	userGroup := api.Group("/users")

	userGroup.Get("/", userHandler.FindAllUsers)
	userGroup.Get("/:id", userHandler.FindUserByID)
	userGroup.Get("/handler/:handler", userHandler.FindUserByHandler)
	userGroup.Get("/email/:email", userHandler.FindUserByEmail)

	postGroup := api.Group("/posts")

	postGroup.Get("/", postHandler.FindPosts)
	postGroup.Get("/:id", postHandler.FindPostByID)

	likeGroup := api.Group("/likes")

	likeGroup.Get("/", likeHandler.FindLikes)
	likeGroup.Get("/:id", likeHandler.FindLikeByID)
	likeGroup.Get("/count", likeHandler.CountLikes)

	commentGroup := api.Group("/comments")

	commentGroup.Get("/", commentHandler.FindComments)
	commentGroup.Get("/:id", commentHandler.FindCommentByID)
}
