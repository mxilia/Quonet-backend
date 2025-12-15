package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/comment/repository"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentUseCase {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *entities.Comment) error {
	if err := s.repo.Save(comment); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) FindComments(authorID uuid.UUID, parentID uuid.UUID, rootID uuid.UUID, page int, limit int) ([]*entities.Comment, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	comments, err := s.repo.Find(authorID, parentID, rootID, offset, limit)
	if err != nil {
		return nil, -1, err
	}

	totalComments, err := s.repo.Count()
	if err != nil {
		return nil, -1, err
	}
	return comments, totalComments, nil
}

func (s *CommentService) FindCommentByID(id uuid.UUID) (*entities.Comment, error) {
	comment, err := s.repo.FindByID(context.TODO(), id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) PatchComment(id uuid.UUID, comment *entities.Comment) error {
	if err := s.repo.Patch(context.TODO(), id, comment); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) DeleteComment(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
