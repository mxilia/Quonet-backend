package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/post/dto"
	"github.com/mxilia/Conflux-backend/internal/post/usecase"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
	"github.com/mxilia/Conflux-backend/pkg/responses"
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

/* No private posts involved */
func (h *HttpPostHandler) FindAllPosts(c *fiber.Ctx) error {
	posts, err := h.usecase.FindAllPosts()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPostsByAuthorID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPostsByAuthorID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPostsByThreadID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPostsByThreadID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
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

func (h *HttpPostHandler) FindPostByTitle(c *fiber.Ctx) error {
	post, err := h.usecase.FindPostByTitle(c.Params("title"))
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponse(post))
}

/* Private posts involved */
func (h *HttpPostHandler) FindAllPostsCoverPrivate(c *fiber.Ctx) error {
	posts, err := h.usecase.FindAllPostsCoverPrivate()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindAllPrivatePosts(c *fiber.Ctx) error {
	posts, err := h.usecase.FindAllPrivatePosts()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPostsCoverPrivateByAuthorID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPostsCoverPrivateByAuthorID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPrivatePostsByAuthorID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPrivatePostsByAuthorID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPostsCoverPrivateByThreadID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPostsCoverPrivateByThreadID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPrivatePostsByThreadID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	posts, err := h.usecase.FindPrivatePostsByThreadID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponseList(posts))
}

func (h *HttpPostHandler) FindPrivatePostByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	post, err := h.usecase.FindPrivatePostByID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToPostResponse(post))
}

func (h *HttpPostHandler) FindPrivatePostByTitle(c *fiber.Ctx) error {
	post, err := h.usecase.FindPrivatePostByTitle(c.Params("title"))
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

	if err := h.usecase.DeletePost(id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
