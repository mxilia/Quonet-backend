package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/thread/repository"
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

func (s *ThreadService) FindThreads(title string, page int, limit int) ([]*entities.Thread, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	threads, err := s.repo.Find(title, offset, limit)
	if err != nil {
		return nil, -1, err
	}

	count, err := s.repo.Count()
	if err != nil {
		return nil, -1, err
	}
	return threads, count, nil
}

func (s *ThreadService) FindThreadByID(id uuid.UUID) (*entities.Thread, error) {
	thread, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (s *ThreadService) DeleteThread(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
