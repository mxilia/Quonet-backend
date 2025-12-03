package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type UserUseCase interface {
	GoogleUserEntry(user *entities.User) (*entities.User, error)
	FindAllUsers() ([]*entities.User, error)
	FindUserByID(id uuid.UUID) (*entities.User, error)
	FindUserByHandler(handler string) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	PatchUser(id uuid.UUID, user *entities.User) error
	DeleteUser(id uuid.UUID) error
}
