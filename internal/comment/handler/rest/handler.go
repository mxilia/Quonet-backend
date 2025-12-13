package handler

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/comment/dto"
	"github.com/mxilia/Quonet-backend/internal/comment/usecase"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

type HttpCommentHandler struct {
	usecase usecase.CommentUseCase
}

func NewHttpCommentHandler(usecase usecase.CommentUseCase) *HttpCommentHandler {
	return &HttpCommentHandler{usecase: usecase}
}

func checkCommentForbidAction(c *fiber.Ctx, h *HttpCommentHandler, commentID uuid.UUID) error {
	existedComment, err := h.usecase.FindCommentByID(commentID)
	if err != nil {
		return err
	}

	if c.Locals("role").(string) == "member" && existedComment.AuthorID != c.Locals("user_id") {
		return appError.ErrForbidden
	}
	return nil
}

func (h *HttpCommentHandler) CreateComment(c *fiber.Ctx) error {
	var req dto.CreateCommentRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.Error(c, err)
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	comment, err := dto.FromCommentCreateRequest(userID, &req)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid create request")
	}

	if err := h.usecase.CreateComment(comment); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create comment")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToCommentResponse(comment))
}

func (h *HttpCommentHandler) FindComments(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 5
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

	parentID := uuid.Nil
	queryParentID := c.Query("parent_id")
	if queryParentID != "" {
		var err error
		parentID, err = uuid.Parse(queryParentID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid parent id")
		}
	}

	rootID := uuid.Nil
	queryRootID := c.Query("root_id")
	if queryRootID != "" {
		var err error
		rootID, err = uuid.Parse(queryRootID)
		if err != nil {
			return responses.ErrorWithMessage(c, err, "invalid root id")
		}
	}

	comments, totalComments, err := h.usecase.FindComments(authorID, parentID, rootID, page, limit)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(fiber.Map{
		"data": dto.ToCommentResponseList(comments),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalComments,
			"totalPages": int(math.Ceil(float64(totalComments) / float64(limit))),
		},
	})
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

	err = checkCommentForbidAction(c, h, id)
	if err != nil {
		return responses.Error(c, err)
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

	err = checkCommentForbidAction(c, h, id)
	if err != nil {
		return responses.Error(c, err)
	}

	if err := h.usecase.DeleteComment(id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
