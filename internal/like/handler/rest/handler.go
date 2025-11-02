package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/like/dto"
	"github.com/mxilia/Conflux-backend/internal/like/usecase"
	"github.com/mxilia/Conflux-backend/pkg/responses"
)

type HttpLikeHandler struct {
	usecase usecase.LikeUseCase
}

func NewHttpLikeHandler(usecase usecase.LikeUseCase) *HttpLikeHandler {
	return &HttpLikeHandler{usecase: usecase}
}

func (h *HttpLikeHandler) CreateLike(c *fiber.Ctx) error {
	var req dto.CreateLikeRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	ownerID, err := uuid.Parse(req.OwnerID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	parentID, err := uuid.Parse(req.ParentID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	like := &entities.Like{OwnerID: ownerID, ParentID: parentID, ParentType: req.ParentType}
	if err := h.usecase.CreateLike(like); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to create like")
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToLikeResponse(like))
}

func (h *HttpLikeHandler) FindAllLikes(c *fiber.Ctx) error {
	likes, err := h.usecase.FindAllLikes()
	if err != nil {
		return responses.Error(c, err)
	}
	return c.JSON(dto.ToLikeResponseList(likes))
}

func (h *HttpLikeHandler) FindLikesByOwnerID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	likes, err := h.usecase.FindLikesByOwnerID(id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToLikeResponseList(likes))
}

func (h *HttpLikeHandler) FindLikesByParentID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	likes, err := h.usecase.FindLikesByParentID(id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToLikeResponseList(likes))
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

func (h *HttpLikeHandler) LikeCountByParentID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	count, err := h.usecase.LikeCountByParentID(c.Params("parent_type"), id)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to get like count")
	}

	return c.JSON(dto.ToLikeCountResponse(count))
}

func (h *HttpLikeHandler) IsParentLikedByMe(c *fiber.Ctx) error {
	parentID, err := uuid.Parse(c.Params("parent_id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	myID, err := uuid.Parse(c.Params("my_id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	isLiked, err := h.usecase.IsParentLikedByMe(c.Params("parent_type"), parentID, myID)
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to get isliked")
	}

	return c.JSON(dto.ToIsLikedResponse(isLiked))
}

func (h *HttpLikeHandler) DeleteLike(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.usecase.DeleteLike(id); err != nil {
		responses.ErrorWithMessage(c, err, "failed to delete like")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
