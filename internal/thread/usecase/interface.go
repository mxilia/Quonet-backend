package usecase

import (
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type ThreadUseCase interface {
	CreateThread(thread *entities.Thread) error
	FindAllThreads() ([]*entities.Thread, error)
	FindThreadByID(id uint) (*entities.Thread, error)
	DeleteThread(id uint) error
}
