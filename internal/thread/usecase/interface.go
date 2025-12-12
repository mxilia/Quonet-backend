package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type ThreadUseCase interface {
	CreateThread(thread *entities.Thread) error
	FindAllThreads(page int, limit int) ([]*entities.Thread, int64, error)
	FindThreadByID(id uuid.UUID) (*entities.Thread, error)
	DeleteThread(id uuid.UUID) error
}
