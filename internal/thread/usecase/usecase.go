package usecase

import (
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/thread/repository"
)

type ThreadService struct {
	repo repository.ThreadRepository
}

func NewThreadService(repo repository.ThreadRepository) ThreadUseCase {
	return &ThreadService{repo: repo}
}

func (s *ThreadService) CreateThread(thread *entities.Thread) error {
	if err := s.repo.Save(thread); err != nil {
		return err
	}
	return nil
}

func (s *ThreadService) GetAllThreads() ([]*entities.Thread, error) {
	threads, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return threads, nil
}

func (s *ThreadService) GetThreadByID(id uint) (*entities.Thread, error) {
	thread, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (s *ThreadService) DeleteThread(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
