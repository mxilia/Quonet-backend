package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/user/repository"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) GoogleUserEntry(user *entities.User) (*entities.User, error) {
	if user.Email == "" {
		return nil, appError.ErrInvalidData
	}

	registeredUser, err := s.repo.FindByEmail(user.Email)
	if err != nil && !errors.Is(err, appError.ErrRecordNotFound) {
		return nil, err
	}
	if registeredUser != nil {
		return registeredUser, nil
	}

	registeredUser = &entities.User{}
	if err := s.repo.Save(registeredUser); err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (s *UserService) FindAllUsers() ([]*entities.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) FindUserByID(id uuid.UUID) (*entities.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindUserByHandler(handler string) (*entities.User, error) {
	user, err := s.repo.FindByHandler(handler)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) PatchUser(id uuid.UUID, user *entities.User) error {
	if err := s.repo.Patch(id, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
