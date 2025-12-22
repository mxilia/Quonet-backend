package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/announcement/repository"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type AnnouncementService struct {
	repo repository.AnnouncementRepository
}

func NewAnnouncementService(repo repository.AnnouncementRepository) AnnouncementUseCase {
	return &AnnouncementService{repo: repo}
}

func (s *AnnouncementService) SaveAnnouncement(announcement *entities.Announcement) error {
	if err := s.repo.Save(announcement); err != nil {
		return err
	}
	return nil
}

func (s *AnnouncementService) FindAnnouncements(page int, limit int) ([]*entities.Announcement, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	announcements, err := s.repo.Find(offset, limit)
	if err != nil {
		return nil, -1, err
	}

	totalAnnouncements, err := s.repo.Count()
	if err != nil {
		return nil, -1, err
	}
	return announcements, totalAnnouncements, nil
}

func (s *AnnouncementService) DeleteAnnouncement(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
