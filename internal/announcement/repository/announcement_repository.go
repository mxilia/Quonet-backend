package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type AnnouncementRepository interface {
	Save(announcement *entities.Announcement) error
	Find(offset int, limit int) ([]*entities.Announcement, error)
	Count() (int64, error)
	Delete(id uuid.UUID) error
}
