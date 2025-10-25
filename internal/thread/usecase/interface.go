package usecase

import (
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type ThreadUseCase interface {
	CreateThread(thread *entities.Thread) error
	GetAllThreads() ([]*entities.Thread, error)
	GetThreadByID(id uint) (*entities.Thread, error)
	DeleteThread(id uint) error
}
