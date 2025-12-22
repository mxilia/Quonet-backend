package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type AnnouncementUseCase interface {
	SaveAnnouncement(announcement *entities.Announcement) error
	FindAnnouncements(page int, limit int) ([]*entities.Announcement, int64, error)
	DeleteAnnouncement(id uuid.UUID) error
}
