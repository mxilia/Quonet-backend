package routes

import (
	"github.com/gofiber/fiber/v2"

	commentHandler "github.com/mxilia/Conflux-backend/internal/comment/handler/rest"
	commentRepository "github.com/mxilia/Conflux-backend/internal/comment/repository"
	commentUseCase "github.com/mxilia/Conflux-backend/internal/comment/usecase"

	likeHandler "github.com/mxilia/Conflux-backend/internal/like/handler/rest"
	likeRepository "github.com/mxilia/Conflux-backend/internal/like/repository"

	likeUseCase "github.com/mxilia/Conflux-backend/internal/like/usecase"
	postHandler "github.com/mxilia/Conflux-backend/internal/post/handler/rest"
	postRepository "github.com/mxilia/Conflux-backend/internal/post/repository"

	postUseCase "github.com/mxilia/Conflux-backend/internal/post/usecase"
	sessionHandler "github.com/mxilia/Conflux-backend/internal/session/handler/rest"
	sessionRepository "github.com/mxilia/Conflux-backend/internal/session/repository"

	sessionUseCase "github.com/mxilia/Conflux-backend/internal/session/usecase"
	threadHandler "github.com/mxilia/Conflux-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Conflux-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Conflux-backend/internal/thread/usecase"

	userHandler "github.com/mxilia/Conflux-backend/internal/user/handler/rest"
	userRepository "github.com/mxilia/Conflux-backend/internal/user/repository"
	userUseCase "github.com/mxilia/Conflux-backend/internal/user/usecase"

	"github.com/mxilia/Conflux-backend/pkg/config"
	"github.com/mxilia/Conflux-backend/pkg/middleware"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {

	/* === Dependencies Wiring === */

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

	likeRepo := likeRepository.NewGormLikeRepository(db)
	likeUseCase := likeUseCase.NewLikeService(likeRepo)
	likeHandler := likeHandler.NewHttpLikeHandler(likeUseCase)

	commentRepo := commentRepository.NewGormCommentRepository(db)
	commentUseCase := commentUseCase.NewCommentService(commentRepo)
	commentHandler := commentHandler.NewHttpCommentHandler(commentUseCase)

	/* === Routes === */

	api := app.Group("/api/v2", middleware.JWTMiddleware(cfg))

	api.Get("/me", userHandler.GetUser)

	threadGroup := api.Group("/threads")

	threadGroup.Post("/", threadHandler.CreateThread)
	threadGroup.Delete("/:id", middleware.RequireAdmin(), threadHandler.DeleteThread)

	sessionGroup := api.Group("/sessions")

	sessionGroup.Patch("/:email", middleware.RequireAdmin(), sessionHandler.RevokeToken)

	userGroup := api.Group("/users")

	userGroup.Patch("/:id", middleware.RequireUser(), userHandler.PatchUser)
	userGroup.Delete("/:id", middleware.RequireUser(), userHandler.DeleteUser)

	postGroup := api.Group("/posts")

	postGroup.Post("/", postHandler.CreatePost)
	postGroup.Get("/all", middleware.RequireUser(), postHandler.FindAllPostsCoverPrivate)
	postGroup.Get("/private", middleware.RequireUser(), postHandler.FindAllPrivatePosts)
	postGroup.Get("/all/author/:id", middleware.RequireUser(), postHandler.FindPostsCoverPrivateByAuthorID)
	postGroup.Get("/private/author/:id", middleware.RequireUser(), postHandler.FindPrivatePostsByAuthorID)
	postGroup.Get("/all/thread/:id", middleware.RequireUser(), postHandler.FindPostsCoverPrivateByThreadID)
	postGroup.Get("/private/thread/:id", middleware.RequireUser(), postHandler.FindPrivatePostsByThreadID)
	postGroup.Get("/private/:id", middleware.RequireUser(), postHandler.FindPrivatePostByID)
	postGroup.Get("/private/title/:title", middleware.RequireUser(), postHandler.FindPrivatePostByTitle)
	postGroup.Patch("/:id", middleware.RequireUser(), postHandler.PatchPost)
	postGroup.Delete("/:id", middleware.RequireUser(), postHandler.DeletePost)

	likeGroup := api.Group("/likes")

	likeGroup.Post("/", likeHandler.CreateLike)
	likeGroup.Get("/", middleware.RequireUser(), likeHandler.FindAllLikes)
	likeGroup.Get("/owner/:id", middleware.RequireUser(), likeHandler.FindLikesByOwnerID)
	likeGroup.Get("/parent/:id", middleware.RequireUser(), likeHandler.FindLikesByParentID)
	likeGroup.Get("/id/:id", middleware.RequireUser(), likeHandler.FindLikeByID)
	likeGroup.Get("/count/:parent_type/:id", middleware.RequireUser(), likeHandler.LikeCountByParentID)
	likeGroup.Get("/is_liked/:parent_type/:parent_id/:my_id", middleware.RequireUser(), likeHandler.IsParentLikedByMe)
	likeGroup.Delete("/:id", middleware.RequireUser(), likeHandler.DeleteLike)

	commentGroup := api.Group("/comments")

	commentGroup.Post("/", commentHandler.CreateComment)
	commentGroup.Get("/:id", middleware.RequireUser(), commentHandler.FindCommentByID)
	commentGroup.Patch("/:id", middleware.RequireUser(), commentHandler.PatchComment)
	commentGroup.Delete("/:id", middleware.RequireUser(), commentHandler.DeleteComment)
}
