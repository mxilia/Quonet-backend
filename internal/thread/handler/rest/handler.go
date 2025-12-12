package rest

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/thread/dto"
	"github.com/mxilia/Quonet-backend/internal/thread/usecase"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

type HttpThreadHandler struct {
	threadUseCase usecase.ThreadUseCase
}

func NewHttpThreadHandler(threadUseCase usecase.ThreadUseCase) *HttpThreadHandler {
	return &HttpThreadHandler{threadUseCase: threadUseCase}
}

func (h *HttpThreadHandler) CreateThread(c *fiber.Ctx) error {
	var req dto.CreateThreadRequest
	if err := c.BodyParser(&req); err != nil {
		return responses.ErrorWithMessage(c, err, "invalid request")
	}

	thread := &entities.Thread{Title: req.Title}
	if err := h.threadUseCase.CreateThread(thread); err != nil {
		return responses.Error(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToThreadResponse(thread))
}

func (h *HttpThreadHandler) FindAllThreads(c *fiber.Ctx) error {
	var (
		page  = c.QueryInt("page", 1)
		limit = 5
	)

	threads, totalThreads, err := h.threadUseCase.FindAllThreads(page, limit)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(fiber.Map{
		"data": dto.ToThreadResponseList(threads),
		"meta": fiber.Map{
			"page":       page,
			"total":      totalThreads,
			"totalPages": int(math.Ceil(float64(totalThreads) / float64(limit))),
		},
	})
}

func (h *HttpThreadHandler) FindThreadByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	thread, err := h.threadUseCase.FindThreadByID(id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToThreadResponse(thread))
}

func (h *HttpThreadHandler) DeleteThread(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return responses.ErrorWithMessage(c, err, "invalid id")
	}

	if err := h.threadUseCase.DeleteThread(id); err != nil {
		return responses.Error(c, err)
	}

	return responses.Message(c, fiber.StatusOK, "thread deleted successfully")
}
