package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	commentUseCase "github.com/mxilia/Quonet-backend/internal/comment/usecase"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/like/repository"
	postUseCase "github.com/mxilia/Quonet-backend/internal/post/usecase"
	"github.com/mxilia/Quonet-backend/internal/transaction"
)

type LikeService struct {
	repo           repository.LikeRepository
	txManager      transaction.TransactionManager
	postUseCase    postUseCase.PostUseCase
	commentUseCase commentUseCase.CommentUseCase
}

func NewLikeService(repo repository.LikeRepository, txManager transaction.TransactionManager, postUseCase postUseCase.PostUseCase, commentUseCase commentUseCase.CommentUseCase) LikeUseCase {
	return &LikeService{
		repo:           repo,
		txManager:      txManager,
		postUseCase:    postUseCase,
		commentUseCase: commentUseCase,
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

func getUpdateLikeCount(s *LikeService, like *entities.Like) (int64, string, uuid.UUID, error) {
	likeds, _, err := s.FindLikes(like.ParentType, like.ParentID, like.OwnerID, 0, 1)
	liked := likeds[0]
	if err != nil {
		return -1, "", uuid.Nil, err
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

		updateCount, status, likedID, err := getUpdateLikeCount(s, like)
		if err != nil {
			return err
		}

		switch like.ParentType {

		case "post":
			post, err := s.postUseCase.FindPostByID(like.ParentID)
			if err != nil {
				return err
			}
			if post == nil {
				return fmt.Errorf("post does not exist")
			}
			if err := s.postUseCase.PatchPost(like.ParentID, &entities.Post{LikeCount: post.LikeCount + updateCount}); err != nil {
				return err
			}

		case "comment":
			comment, err := s.commentUseCase.FindCommentByID(like.ParentID)
			if err != nil {
				return err
			}
			if comment == nil {
				return fmt.Errorf("comment does not exist")
			}
			if err := s.commentUseCase.PatchComment(like.ParentID, &entities.Comment{LikeCount: comment.LikeCount + updateCount}); err != nil {
				return err
			}

		default:
			return fmt.Errorf("invalid parent type")
		}

		if err := s.repo.Delete(txCtx, likedID); err != nil {
			return err
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

	likes, err := s.repo.Find(parentType, ownerID, parentID, offset, limit)
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
