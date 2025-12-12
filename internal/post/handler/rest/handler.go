package rest

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/post/dto"
	"github.com/mxilia/Quonet-backend/internal/post/usecase"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/responses"
	"github.com/mxilia/Quonet-backend/utils/format"
)

type HttpPostHandler struct {
	usecase usecase.PostUseCase
}

func NewHttpPostHandler(usecase usecase.PostUseCase) *HttpPostHandler {
	return &HttpPostHandler{usecase: usecase}
}

func (h *HttpPostHandler) CreatePost(c *fiber.Ctx) error {
	var req dto.CreatePostRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	authorID, err := uuid.Parse(req.AuthorID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid author id")
	}

	threadID, err := uuid.Parse(req.ThreadID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid author id")
	}

	post := &entities.Post{Title: req.Title, AuthorID: authorID, ThreadID: threadID, Content: req.Content, ThumbnailUrl: req.ThumbnailUrl}
	if err := h.usecase.CreatePost(post); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create post")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToPostResponse(post))
}

func checkPostForbidAction(c *fiber.Ctx, h *HttpPostHandler, postID uuid.UUID) error {
	userID := c.Locals("user_id").(string)
	if userID == "" {
		return appError.ErrUnauthorized
	}

	authorID, err := uuid.Parse(userID)
	if err != nil {
		return appError.ErrUnauthorized
	}

	existedPost, err := h.usecase.FindPrivatePostByID(postID)
	if err != nil {
		return err
	}

	if c.Locals("role").(string) == "member" && existedPost.AuthorID != authorID {
		return appError.ErrForbidden
	}
	return nil
}

/* No private posts involved */
func (h *HttpPostHandler) FindPosts(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 3
	)

	authorID := uuid.Nil
	queryAuthorID := c.Query("author_id")
	if queryAuthorID != "" {
		var err error
		authorID, err = uuid.Parse(queryAuthorID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid author id")
		}
	}

	threadID := uuid.Nil
	queryThreadID := c.Query("thread_id")
	if queryThreadID != "" {
		var err error
		threadID, err = uuid.Parse(queryThreadID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid thread id")
		}
	}

	title := c.Query("title")
	if title != "" {
		title = format.DashToSpace(title)
	}

	posts, totalPosts, err := h.usecase.FindPosts(authorID, threadID, title, page, limit)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(fiber.Map{
		"data": dto.ToPostResponseList(posts),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalPosts,
			"totalPages": int(math.Ceil(float64(totalPosts) / float64(limit))),
		},
	})
}

func (h *HttpPostHandler) FindPostByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	post, err := h.usecase.FindPostByID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponse(post))
}

/* Private posts involved */
func (h *HttpPostHandler) FindPrivatePosts(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 3
	)

	authorID := uuid.Nil
	queryAuthorID := c.Query("author_id")
	if queryAuthorID != "" {
		var err error
		authorID, err = uuid.Parse(queryAuthorID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid author id")
		}
	}

	if queryAuthorID != "" && c.Locals("role").(string) == "member" {
		userID := c.Locals("user_id").(string)
		if userID == "" {
			return responses.Error(c, appError.ErrUnauthorized)
		}

		var err error
		authorID, err = uuid.Parse(userID)
		if err != nil {
			return responses.Error(c, appError.ErrUnauthorized)
		}
	}

	threadID := uuid.Nil
	queryThreadID := c.Query("thread_id")
	if queryThreadID != "" {
		var err error
		threadID, err = uuid.Parse(queryThreadID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid thread id")
		}
	}

	title := c.Query("title")
	if title != "" {
		title = format.DashToSpace(title)
	}

	posts, totalPosts, err := h.usecase.FindPrivatePosts(authorID, threadID, title, page, limit)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(fiber.Map{
		"data": dto.ToPostResponseList(posts),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalPosts,
			"totalPages": int(math.Ceil(float64(totalPosts) / float64(limit))),
		},
	})
}

func (h *HttpPostHandler) FindPrivatePostByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	err = checkPostForbidAction(c, h, id)
	if err != nil {
		return responses.Error(c, err)
	}

	post, err := h.usecase.FindPrivatePostByID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponse(post))
}

func (h *HttpPostHandler) PatchPost(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	err = checkPostForbidAction(c, h, id)
	if err != nil {
		return responses.Error(c, err)
	}

	var req dto.PostPatchRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, err)
	}

	postInfo := dto.FromPostPatchRequest(&req)
	if err := h.usecase.PatchPost(id, postInfo); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "patch successfully")
}

func (h *HttpPostHandler) DeletePost(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	err = checkPostForbidAction(c, h, id)
	if err != nil {
		return responses.Error(c, err)
	}

	if err := h.usecase.DeletePost(id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
