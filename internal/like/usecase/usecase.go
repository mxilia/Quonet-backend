package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/like/repository"
)

type LikeService struct {
	repo repository.LikeRepository
}

func NewLikeService(repo repository.LikeRepository) LikeUseCase {
	return &LikeService{repo: repo}
}

func (s *LikeService) CreateLike(like *entities.Like) error {
	if err := s.repo.Save(like); err != nil {
		return err
	}
	return nil
}

func (s *LikeService) FindAllLikes() ([]*entities.Like, error) {
	likes, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (s *LikeService) FindLikesByOwnerID(id uuid.UUID) ([]*entities.Like, error) {
	likes, err := s.repo.FindByOwnerID(id)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (s *LikeService) FindLikesByParentID(id uuid.UUID) ([]*entities.Like, error) {
	likes, err := s.repo.FindByParentID(id)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

func (s *LikeService) FindLikeByID(id uuid.UUID) (*entities.Like, error) {
	like, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (s *LikeService) LikeCountByParentID(parentType string, id uuid.UUID) (int64, error) {
	count, err := s.repo.CountByParentID(parentType, id)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *LikeService) IsParentLikedByMe(parentType string, parentID uuid.UUID, myID uuid.UUID) (bool, error) {
	isLiked, err := s.repo.IsParentLikedByMe(parentType, parentID, myID)
	if err != nil {
		return false, err
	}
	return isLiked, nil
}

func (s *LikeService) DeleteLike(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
