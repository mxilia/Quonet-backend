package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	commentRepository "github.com/mxilia/Quonet-backend/internal/comment/repository"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/like/repository"
	postRepository "github.com/mxilia/Quonet-backend/internal/post/repository"
	"github.com/mxilia/Quonet-backend/internal/transaction"
)

type LikeService struct {
	repo        repository.LikeRepository
	txManager   transaction.TransactionManager
	postRepo    postRepository.PostRepository
	commentRepo commentRepository.CommentRepository
}

func NewLikeService(repo repository.LikeRepository, txManager transaction.TransactionManager, postRepository postRepository.PostRepository, commentRepository commentRepository.CommentRepository) LikeUseCase {
	return &LikeService{
		repo:        repo,
		txManager:   txManager,
		postRepo:    postRepository,
		commentRepo: commentRepository,
	}
}

func calcUpdateCount(oldPositive bool, newPositive bool) (int64, string) {
	switch {
	case oldPositive && newPositive:
		return -1, ""
	case oldPositive && !newPositive:
		return -2, "create"
	case !oldPositive && newPositive:
		return 2, "create"
	default:
		return 1, ""
	}
}

func getUpdateLikeCount(s *LikeService, like *entities.Like, txCtx context.Context) (int64, string, uuid.UUID, error) {
	likeds, err := s.repo.Find(txCtx, like.ParentType, like.OwnerID, like.ParentID, 0, 1)

	if err != nil {
		return -1, "", uuid.Nil, err
	}

	var liked *entities.Like
	if len(likeds) > 0 {
		liked = likeds[0]
	} else {
		liked = nil
	}

	if liked == nil {
		if like.IsPositive {
			return 1, "create", uuid.Nil, nil
		} else {
			return -1, "create", uuid.Nil, nil
		}
	}

	updateCount, status := calcUpdateCount(liked.IsPositive, like.IsPositive)
	return updateCount, status, liked.ID, nil
}

func (s *LikeService) CreateLike(ctx context.Context, like *entities.Like) error {
	return s.txManager.Do(ctx, func(txCtx context.Context) error {

		updateCount, status, likedID, err := getUpdateLikeCount(s, like, txCtx)
		if err != nil {
			return err
		}

		switch like.ParentType {

		case "post":
			post, err := s.postRepo.FindByID(txCtx, like.ParentID)
			if err != nil {
				return err
			}
			if post == nil {
				return fmt.Errorf("post does not exist")
			}
			if err := s.postRepo.Patch(txCtx, like.ParentID, &entities.Post{LikeCount: post.LikeCount + updateCount}); err != nil {
				return err
			}

		case "comment":
			comment, err := s.commentRepo.FindByID(txCtx, like.ParentID)
			if err != nil {
				return err
			}
			if comment == nil {
				return fmt.Errorf("comment does not exist")
			}
			if err := s.commentRepo.Patch(txCtx, like.ParentID, &entities.Comment{LikeCount: comment.LikeCount + updateCount}); err != nil {
				return err
			}

		default:
			return fmt.Errorf("invalid parent type")
		}

		fmt.Println("liked id:", likedID)
		if likedID != uuid.Nil {
			if err := s.repo.Delete(txCtx, likedID); err != nil {
				return err
			}
		}
		if status == "create" {
			if err := s.repo.Save(txCtx, like); err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *LikeService) FindLikes(parentType string, ownerID uuid.UUID, parentID uuid.UUID, page int, limit int) ([]*entities.Like, int64, error) {
	if page < 1 {
		page = 1
	}

	offset := (page - 1) * limit

	likes, err := s.repo.Find(context.TODO(), parentType, ownerID, parentID, offset, limit)
	if err != nil {
		return nil, -1, err
	}

	totalLikes, err := s.repo.Count(parentType, ownerID, parentID)
	if err != nil {
		return nil, -1, err
	}
	return likes, totalLikes, nil
}

func (s *LikeService) FindLikeByID(id uuid.UUID) (*entities.Like, error) {
	like, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return like, nil
}

func (s *LikeService) CountLikes(parentType string, ownerID uuid.UUID, parentID uuid.UUID) (int64, error) {
	count, err := s.repo.Count(parentType, ownerID, parentID)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (s *LikeService) DeleteLike(id uuid.UUID) error {
	if err := s.repo.Delete(context.TODO(), id); err != nil {
		return err
	}
	return nil
}
