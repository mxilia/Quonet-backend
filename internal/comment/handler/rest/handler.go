package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/comment/dto"
	"github.com/mxilia/Conflux-backend/internal/comment/usecase"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
	"github.com/mxilia/Conflux-backend/pkg/responses"
)

type HttpCommentHandler struct {
	usecase usecase.CommentUseCase
}

func NewHttpCommentHandler(usecase usecase.CommentUseCase) *HttpCommentHandler {
	return &HttpCommentHandler{usecase: usecase}
}

func (h *HttpCommentHandler) CreateComment(c *fiber.Ctx) error {
	var req dto.CreateCommentRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, err)
	}

	comment, err := dto.FromCommentCreateRequest(&req)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid create request")
	}

	if err := h.usecase.CreateComment(comment); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create comment")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToCommentResponse(comment))
}

func (h *HttpCommentHandler) FindAllComments(c *fiber.Ctx) error {
	comments, err := h.usecase.FindAllComments()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCommentResponseList(comments))
}

func (h *HttpCommentHandler) FindCommentsByAuthorID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	comments, err := h.usecase.FindCommentsByAuthorID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCommentResponseList(comments))
}

func (h *HttpCommentHandler) FindCommentsByParentID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	comments, err := h.usecase.FindCommentsByParentID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCommentResponseList(comments))
}

func (h *HttpCommentHandler) FindCommentsByRootID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	comments, err := h.usecase.FindCommentsByRootID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCommentResponseList(comments))
}

func (h *HttpCommentHandler) FindCommentByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	comment, err := h.usecase.FindCommentByID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToCommentResponse(comment))
}

func (h *HttpCommentHandler) PatchComment(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	var req dto.CommentPatchRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, err)
	}

	comment := dto.FromCommentPatchRequest(&req)
	if err := h.usecase.PatchComment(id, comment); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "patch successfully")
}

func (h *HttpCommentHandler) DeleteComment(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	if err := h.usecase.DeleteComment(id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
