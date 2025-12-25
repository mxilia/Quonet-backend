package usecase

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/post/repository"
	"github.com/mxilia/Quonet-backend/internal/transaction"
	"github.com/mxilia/Quonet-backend/pkg/database"
)

type PostService struct {
	repo           repository.PostRepository
	storageService *database.StorageService
	txManager      transaction.TransactionManager
}

func NewPostService(repo repository.PostRepository, storageService *database.StorageService, txManager transaction.TransactionManager) PostUseCase {
	return &PostService{
		repo:           repo,
		storageService: storageService,
		txManager:      txManager,
	}
}

func (s *PostService) CreatePost(ctx context.Context, post *entities.Post, file io.Reader, filename string, contentType string) error {
	if file != nil {
		return s.txManager.Do(ctx, func(txCtx context.Context) error {

			var (
				ext        = filepath.Ext(filename)
				uploadPath = fmt.Sprintf("thumbnails/%s%s", uuid.New().String(), ext)
			)

			post.ThumbnailUrl = uploadPath

			if err := s.repo.Save(txCtx, post); err != nil {
				return err
			}
			fmt.Println("before supa")
			if err := s.storageService.UploadFile(uploadPath, file, contentType); err != nil {
				return err
			}
			return nil
		})
	}

	return s.repo.Save(ctx, post)
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
	post, err := s.repo.FindByID(context.TODO(), id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FindTopLikedPosts(authorID uuid.UUID, threadID uuid.UUID, title string, limit int) ([]*entities.Post, error) {
	if limit < 0 {
		limit = 0
	}

	posts, err := s.repo.FindTopLiked(authorID, threadID, title, limit)
	if err != nil {
		return nil, err
	}
	return posts, nil
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

func (s *PostService) FindNoFilterPostByID(id uuid.UUID) (*entities.Post, error) {
	post, err := s.repo.FindNoFilterByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) PatchPost(id uuid.UUID, post *entities.Post) error {
	if err := s.repo.Patch(context.TODO(), id, post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(ctx context.Context, id uuid.UUID) error {
	post, err := s.repo.FindNoFilterByID(id)
	if err != nil {
		return err
	}

	if post.ThumbnailUrl != "" {
		return s.txManager.Do(ctx, func(txCtx context.Context) error {
			if err := s.repo.Delete(ctx, id); err != nil {
				return err
			}

			if err := s.storageService.DeleteFile(post.ThumbnailUrl); err != nil {
				return err
			}
			return nil
		})
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
