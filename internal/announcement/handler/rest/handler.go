package rest

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/announcement/dto"
	"github.com/mxilia/Quonet-backend/internal/announcement/usecase"
	"github.com/mxilia/Quonet-backend/internal/entities"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

type HttpAnnouncementHandler struct {
	usecase usecase.AnnouncementUseCase
}

func NewHttpAnnouncementHandler(usecase usecase.AnnouncementUseCase) *HttpAnnouncementHandler {
	return &HttpAnnouncementHandler{usecase: usecase}
}

func (h *HttpAnnouncementHandler) SaveAnnouncement(c *fiber.Ctx) error {
	var req dto.CreateAnnouncementRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	userID, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return responses.Error(c, appError.ErrUnauthorized)
	}

	announcement := &entities.Announcement{AuthorID: userID, Content: req.Content}
	if err := h.usecase.SaveAnnouncement(announcement); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToAnnouncementResponse(announcement))
}

func (h *HttpAnnouncementHandler) FindAnnouncements(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 5
	)

	announcements, totalAnnouncements, err := h.usecase.FindAnnouncements(page, limit)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(fiber.Map{
		"data": dto.ToAnnouncementResponseList(announcements),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalAnnouncements,
			"totalPages": int(math.Ceil(float64(totalAnnouncements) / float64(limit))),
		},
	})
}

func (h *HttpAnnouncementHandler) DeleteAnnouncement(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "invalid id")
	}

	if err := h.usecase.DeleteAnnouncement(id); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to delete")
	}
	return responses.Message(c, fiber.StatusOK, "deleted successfully")
}
