package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/post/repository"
)

type PostService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostUseCase {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *entities.Post) error {
	if err := s.repo.Save(post); err != nil {
		return err
	}
	return nil
}

/* No private posts involved */
func (s *PostService) FindPosts(authorID uuid.UUID, threadID uuid.UUID, title string, page int, limit int) ([]*entities.Post, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	posts, err := s.repo.Find(authorID, threadID, title, offset, limit)
	if err != nil {
		return nil, -1, err
	}

	totalPosts, err := s.repo.Count(false, authorID, threadID, title)
	if err != nil {
		return nil, -1, err
	}

	return posts, totalPosts, nil
}

func (s *PostService) FindPostByID(id uuid.UUID) (*entities.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

/* Private posts involved */
func (s *PostService) FindPrivatePosts(authorID uuid.UUID, threadID uuid.UUID, title string, page int, limit int) ([]*entities.Post, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	posts, err := s.repo.FindPrivate(authorID, threadID, title, offset, limit)
	if err != nil {
		return nil, -1, err
	}

	totalPosts, err := s.repo.Count(true, authorID, threadID, title)
	if err != nil {
		return nil, -1, err
	}

	return posts, totalPosts, nil
}

func (s *PostService) FindPrivatePostByID(id uuid.UUID) (*entities.Post, error) {
	post, err := s.repo.FindPrivateByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) PatchPost(id uuid.UUID, post *entities.Post) error {
	if err := s.repo.Patch(id, post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
