package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/internal/post/repository"
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
func (s *PostService) FindAllPosts() ([]*entities.Post, error) {
	posts, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostsByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostsByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindByThreadID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostByID(id uuid.UUID) (*entities.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FindPostByTitle(title string) (*entities.Post, error) {
	post, err := s.repo.FindByTitle(title)
	if err != nil {
		return nil, err
	}
	return post, nil
}

/* Private posts involved */
func (s *PostService) FindAllPostsCoverPrivate() ([]*entities.Post, error) {
	posts, err := s.repo.FindAllCoverPrivate()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindAllPrivatePosts() ([]*entities.Post, error) {
	posts, err := s.repo.FindAllPrivate()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostsCoverPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindCoverPrivateByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPrivatePostsByAuthorID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindPrivateByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostsCoverPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindCoverPrivateByThreadID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPrivatePostsByThreadID(id uuid.UUID) ([]*entities.Post, error) {
	posts, err := s.repo.FindPrivateByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPrivatePostByID(id uuid.UUID) (*entities.Post, error) {
	post, err := s.repo.FindPrivateByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FindPrivatePostByTitle(title string) (*entities.Post, error) {
	post, err := s.repo.FindPrivateByTitle(title)
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
