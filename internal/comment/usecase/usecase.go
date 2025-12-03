package usecase

import (
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

func (s *CommentService) FindAllComments() ([]*entities.Comment, error) {
	comments, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentsByAuthorID(id uuid.UUID) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentsByParentID(id uuid.UUID) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByParentID(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentsByRootID(id uuid.UUID) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByRootID(id)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentByID(id uuid.UUID) (*entities.Comment, error) {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) PatchComment(id uuid.UUID, comment *entities.Comment) error {
	if err := s.repo.Patch(id, comment); err != nil {
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
