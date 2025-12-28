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

	announcementHandler "github.com/mxilia/Quonet-backend/internal/announcement/handler/rest"
	announcementRepository "github.com/mxilia/Quonet-backend/internal/announcement/repository"
	announcementUseCase "github.com/mxilia/Quonet-backend/internal/announcement/usecase"

	"github.com/mxilia/Quonet-backend/pkg/config"
	"github.com/mxilia/Quonet-backend/pkg/database"
	"github.com/mxilia/Quonet-backend/pkg/middleware"
	ratelimit "github.com/mxilia/Quonet-backend/pkg/middleware/rate_limit"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app *fiber.App, db *gorm.DB, storageService *database.StorageService, limiter *ratelimit.RateLimiter, cfg *config.Config) {

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
	postUseCase := postUseCase.NewPostService(postRepo, storageService, txManager)
	postHandler := postHandler.NewHttpPostHandler(postUseCase, storageService)

	commentRepo := commentRepository.NewGormCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentService(commentRepo)
	commentHandler := commentHandler.NewHttpCommentHandler(commentUseCase)

	likeRepo := likeRepository.NewGormLikeRepository(db)
	likeUseCase := likeUseCase.NewLikeService(likeRepo, txManager, postRepo, commentRepo)
	likeHandler := likeHandler.NewHttpLikeHandler(likeUseCase)

	announcementRepo := announcementRepository.NewGormAnnouncementRepository(db)
	announcementUseCase := announcementUseCase.NewAnnouncementService(announcementRepo)
	announcementHandler := announcementHandler.NewHttpAnnouncementHandler(announcementUseCase)

	/* === Routes === */

	api := app.Group("/api/v2", middleware.JWTMiddleware(cfg, sessionUseCase, userUseCase))

	api.Get(
		"/me",
		limiter.Use(ratelimit.UserRead, ratelimit.UserKey),
		userHandler.GetUser,
	)

	authGroup := api.Group("/auth")

	authGroup.Post(
		"/logout",
		limiter.Use(ratelimit.AuthLogout, ratelimit.UserKey),
		sessionHandler.Logout,
	)

	threadGroup := api.Group("/threads")

	threadGroup.Post(
		"/",
		middleware.RequireAdmin(),
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		threadHandler.CreateThread,
	)
	threadGroup.Delete(
		"/:id",
		middleware.RequireAdmin(),
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		threadHandler.DeleteThread,
	)

	sessionGroup := api.Group("/sessions")

	sessionGroup.Patch(
		"/:email",
		middleware.RequireAdmin(),
		limiter.Use(ratelimit.AdminWrite, ratelimit.UserKey),
		sessionHandler.RevokeToken,
	)

	userGroup := api.Group("/users")

	userGroup.Patch(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		userHandler.PatchUser,
	)
	userGroup.Delete(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		userHandler.DeleteUser,
	)

	postGroup := api.Group("/posts")

	postGroup.Post(
		"/",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		postHandler.CreatePost,
	)
	postGroup.Get(
		"/private",
		limiter.Use(ratelimit.UserRead, ratelimit.UserKey),
		postHandler.FindPrivatePosts,
	)
	postGroup.Get(
		"/private/:id",
		limiter.Use(ratelimit.UserRead, ratelimit.UserKey),
		postHandler.FindPrivatePostByID,
	)
	postGroup.Patch(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		postHandler.PatchPost,
	)
	postGroup.Delete(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		postHandler.DeletePost,
	)

	likeGroup := api.Group("/likes")

	likeGroup.Post(
		"/",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		likeHandler.CreateLike,
	)
	likeGroup.Get(
		"/attributes/state",
		limiter.Use(ratelimit.UserRead, ratelimit.UserKey),
		likeHandler.GetLikeState,
	)
	// likeGroup.Delete("/:id", middleware.RequireUser(), likeHandler.DeleteLike)

	commentGroup := api.Group("/comments")

	commentGroup.Post(
		"/",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		commentHandler.CreateComment,
	)
	commentGroup.Patch(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		commentHandler.PatchComment,
	)
	commentGroup.Delete(
		"/:id",
		limiter.Use(ratelimit.UserWrite, ratelimit.UserKey),
		commentHandler.DeleteComment,
	)

	announcementGroup := api.Group("/announcements")

	announcementGroup.Post(
		"/",
		middleware.RequireAdmin(),
		limiter.Use(ratelimit.AdminWrite, ratelimit.UserKey),
		announcementHandler.SaveAnnouncement,
	)
	announcementGroup.Delete(
		"/:id",
		middleware.RequireAdmin(),
		limiter.Use(ratelimit.AdminWrite, ratelimit.UserKey),
		announcementHandler.DeleteAnnouncement,
	)
}
