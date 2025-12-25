package rest

import (
	"io"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/post/dto"
	"github.com/mxilia/Quonet-backend/internal/post/usecase"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/database"
	"github.com/mxilia/Quonet-backend/pkg/responses"
	"github.com/mxilia/Quonet-backend/utils/format"
)

type HttpPostHandler struct {
	usecase        usecase.PostUseCase
	storageService *database.StorageService
}

func NewHttpPostHandler(usecase usecase.PostUseCase, storageService *database.StorageService) *HttpPostHandler {
	return &HttpPostHandler{
		usecase:        usecase,
		storageService: storageService,
	}
}

func (h *HttpPostHandler) CreatePost(c *fiber.Ctx) error {
	req := &dto.CreatePostRequest{
		Title:    c.FormValue("title"),
		ThreadID: c.FormValue("thread_id"),
		Content:  c.FormValue("content"),
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	threadID, err := uuid.Parse(req.ThreadID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid author id")
	}

	fileHeader, err := c.FormFile("thumbnail")

	var (
		fileName    = ""
		contentType = ""
	)

	var file io.Reader
	if fileHeader != nil {
		reader, err := fileHeader.Open()
		if err != nil {
			return responses.Error(c, appError.ErrInternalServer)
		}
		defer reader.Close()

		const MaxSize = 1 << 20
		if fileHeader.Size > MaxSize {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "file size must not exceed 1MB")
		}

		allowedTypes := map[string]bool{
			"image/png":  true,
			"image/jpeg": true,
			"image/webp": true,
		}
		if !allowedTypes[fileHeader.Header.Get("Content-Type")] {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid file type")
		}

		file = reader
		fileName = fileHeader.Filename
		contentType = fileHeader.Header.Get("Content-Type")
	}

	post := &entities.Post{Title: req.Title, AuthorID: userID, ThreadID: threadID, Content: req.Content, ThumbnailUrl: ""}
	if err := h.usecase.CreatePost(c.Context(), post, file, fileName, contentType); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create post")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToPostResponse(post, h.storageService))
}

func checkPostForbidAction(c *fiber.Ctx, h *HttpPostHandler, postID uuid.UUID) error {
	post, err := h.usecase.FindNoFilterPostByID(postID)
	if err != nil {
		return err
	}

	if c.Locals("role").(string) == "member" && post.AuthorID != c.Locals("user_id") {
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
		"data": dto.ToPostResponseList(posts, h.storageService),
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
	return c.JSON(dto.ToPostResponse(post, h.storageService))
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
		"data": dto.ToPostResponseList(posts, h.storageService),
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
	return c.JSON(dto.ToPostResponse(post, h.storageService))
}

func (h *HttpPostHandler) FindTopLikedPosts(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 3)

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

	posts, err := h.usecase.FindTopLikedPosts(authorID, threadID, title, limit)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts, h.storageService))
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

	if err := h.usecase.DeletePost(c.Context(), id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
