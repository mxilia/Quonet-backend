package rest

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/like/dto"
	"github.com/mxilia/Quonet-backend/internal/like/usecase"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

type HttpLikeHandler struct {
	usecase usecase.LikeUseCase
}

func NewHttpLikeHandler(usecase usecase.LikeUseCase) *HttpLikeHandler {
	return &HttpLikeHandler{usecase: usecase}
}

func checkLikeForbidAction(c *fiber.Ctx, h *HttpLikeHandler, likeID uuid.UUID) error {
	existedLike, err := h.usecase.FindLikeByID(likeID)
	if err != nil {
		return err
	}

	if c.Locals("role").(string) == "member" && existedLike.OwnerID != c.Locals("user_id") {
		return appError.ErrForbidden
	}
	return nil
}

func (h *HttpLikeHandler) CreateLike(c *fiber.Ctx) error {
	var req dto.CreateLikeRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	parentID, err := uuid.Parse(req.ParentID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	like := &entities.Like{OwnerID: userID, ParentID: parentID, ParentType: req.ParentType, IsPositive: req.IsPositive}
	if err := h.usecase.CreateLike(c.Context(), like); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create like")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToLikeResponse(like))
}

func (h *HttpLikeHandler) FindLikes(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 15
	)

	parentType := c.Query("parent_type")
	if parentType != "comment" && parentType != "post" && parentType != "" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	ownerID := uuid.Nil
	unparsedOwnerID := c.Query("owner_id")
	if unparsedOwnerID != "" {
		var err error
		ownerID, err = uuid.Parse(unparsedOwnerID)
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid owner id")
		}
	}

	parentID := uuid.Nil
	unparsedParentID := c.Query("parent_id")
	if unparsedParentID != "" {
		var err error
		parentID, err = uuid.Parse(unparsedParentID)
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid parent id")
		}
	}

	likes, totalLikes, err := h.usecase.FindLikes(parentType, ownerID, parentID, page, limit)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(fiber.Map{
		"data": dto.ToLikeResponseList(likes),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalLikes,
			"totalPages": int(math.Ceil(float64(totalLikes) / float64(limit))),
		},
	})
}

func (h *HttpLikeHandler) FindLikeByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	like, err := h.usecase.FindLikeByID(id)
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToLikeResponse(like))
}

func (h *HttpLikeHandler) GetLikeState(c *fiber.Ctx) error {
	ownerID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	parentType := c.Query("parent_type")
	if parentType != "comment" && parentType != "post" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	unparsedParentID := c.Query("parent_id")
	if unparsedParentID == "" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	parentID, err := uuid.Parse(unparsedParentID)
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid parent id")
	}

	likes, _, err := h.usecase.FindLikes(parentType, ownerID, parentID, 1, 1)
	if err != nil {
		return responses.Error(c, err)
	}

	var (
		IsLikePositive = false
		IsLiked        = false
	)

	if len(likes) > 0 {
		IsLiked = true
		IsLikePositive = likes[0].IsPositive
	}

	return c.JSON(dto.ToLikeStateResponse(IsLiked, IsLikePositive))
}

func (h *HttpLikeHandler) CountLikes(c *fiber.Ctx) error {
	parentType := c.Query("parent_type")
	if parentType != "comment" && parentType != "post" && parentType != "" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	ownerID := uuid.Nil
	unparsedOwnerID := c.Query("owner_id")
	if unparsedOwnerID != "" {
		var err error
		ownerID, err = uuid.Parse(unparsedOwnerID)
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid owner id")
		}
	}

	parentID := uuid.Nil
	unparsedParentID := c.Query("parent_id")
	if unparsedParentID != "" {
		var err error
		parentID, err = uuid.Parse(unparsedParentID)
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid parent id")
		}
	}

	count, err := h.usecase.CountLikes(parentType, ownerID, parentID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to get like count")
	}

	return c.JSON(dto.ToLikeCountResponse(count))
}

func (h *HttpLikeHandler) DeleteLike(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	checkLikeForbidAction(c, h, id)

	if err := h.usecase.DeleteLike(id); err != nil {
		responses.ErrorWithMessage(c, err, "failed to delete like")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
