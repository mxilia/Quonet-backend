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
	threadHandler "github.com/mxilia/Conflux-backend/internal/thread/handler/rest"
	threadRepository "github.com/mxilia/Conflux-backend/internal/thread/repository"
	threadUseCase "github.com/mxilia/Conflux-backend/internal/thread/usecase"
	userHandler "github.com/mxilia/Conflux-backend/internal/user/handler/rest"
	userRepository "github.com/mxilia/Conflux-backend/internal/user/repository"
	userUseCase "github.com/mxilia/Conflux-backend/internal/user/usecase"
	"github.com/mxilia/Conflux-backend/pkg/config"
	"gorm.io/gorm"
)

func RegisterPrivateRoutes(app *fiber.App, db *gorm.DB, cfg *config.Config) {

	/* === Dependencies Wiring === */

	threadRepo := threadRepository.NewGormThreadRepository(db)
	threadUseCase := threadUseCase.NewThreadService(threadRepo)
	threadHandler := threadHandler.NewHttpThreadHandler(threadUseCase)

	userRepo := userRepository.NewGormUserRepository(db)
	userUseCase := userUseCase.NewUserService(userRepo)
	userHandler := userHandler.NewHttpUserHandler(userUseCase, cfg)

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

	api := app.Group("/api/v2")

	threadGroup := api.Group("/threads")

	threadGroup.Post("/", threadHandler.CreateThread)
	threadGroup.Delete("/:id", threadHandler.DeleteThread)

	userGroup := api.Group("/users")

	userGroup.Patch("/:id", userHandler.PatchUser)
	userGroup.Delete("/:id", userHandler.DeleteUser)

	postGroup := api.Group("/posts")

	postGroup.Post("/", postHandler.CreatePost)
	postGroup.Get("/all", postHandler.FindAllPostsCoverPrivate)
	postGroup.Get("/private", postHandler.FindAllPrivatePosts)
	postGroup.Get("/all/author/:id", postHandler.FindPostsCoverPrivateByAuthorID)
	postGroup.Get("/private/author/:id", postHandler.FindPrivatePostsByAuthorID)
	postGroup.Get("/all/thread/:id", postHandler.FindPostsCoverPrivateByThreadID)
	postGroup.Get("/private/thread/:id", postHandler.FindPrivatePostsByThreadID)
	postGroup.Get("/private/:id", postHandler.FindPrivatePostByID)
	postGroup.Get("/private/title/:title", postHandler.FindPrivatePostByTitle)
	postGroup.Patch("/:id", postHandler.PatchPost)
	postGroup.Delete("/:id", postHandler.DeletePost)

	likeGroup := api.Group("/likes")

	likeGroup.Post("/", likeHandler.CreateLike)
	likeGroup.Get("/", likeHandler.FindAllLikes)
	likeGroup.Get("/owner/:id", likeHandler.FindLikesByOwnerID)
	likeGroup.Get("/parent/:id", likeHandler.FindLikesByParentID)
	likeGroup.Get("/id/:id", likeHandler.FindLikeByID)
	likeGroup.Get("/count/:parent_type/:id", likeHandler.LikeCountByParentID)
	likeGroup.Get("/is_liked/:parent_type/:parent_id/:my_id", likeHandler.IsParentLikedByMe)
	likeGroup.Delete("/:id", likeHandler.DeleteLike)

	commentGroup := api.Group("/comments")

	commentGroup.Post("/", commentHandler.CreateComment)
	commentGroup.Get("/:id", commentHandler.FindCommentByID)
	commentGroup.Patch("/:id", commentHandler.PatchComment)
	commentGroup.Delete("/:id", commentHandler.DeleteComment)
}
