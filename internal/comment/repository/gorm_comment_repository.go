package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/transaction"
	"gorm.io/gorm"
)

type GormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) CommentRepository {
	return &GormCommentRepository{db: db}
}

func (r *GormCommentRepository) Save(comment *entities.Comment) error {
	if err := r.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormCommentRepository) Find(authorID uuid.UUID, parentID uuid.UUID, rootID uuid.UUID, offset int, limit int) ([]*entities.Comment, error) {
	query := r.db.Model(&entities.Comment{})
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if parentID != uuid.Nil {
		query = query.Where("parent_id = ?", parentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}

	if rootID != uuid.Nil {
		query = query.Where("root_id = ?", rootID)
	}

	var commentsValue []entities.Comment
	if err := query.Preload("Author").Limit(limit).Offset(offset).Find(&commentsValue).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentsValue))
	for i := range commentsValue {
		comments[i] = &commentsValue[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Comment, error) {
	tx := transaction.GetTx(ctx, r.db)

	var comment entities.Comment
	if err := tx.Preload("Author").Where("id = ?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *GormCommentRepository) Count(authorID uuid.UUID, parentID uuid.UUID, rootID uuid.UUID) (int64, error) {
	query := r.db.Model(&entities.Comment{})
	if authorID != uuid.Nil {
		query = query.Where("author_id = ?", authorID)
	}

	if parentID != uuid.Nil {
		query = query.Where("parent_id = ?", parentID)
	} else {
		query = query.Where("parent_id IS NULL")
	}

	if rootID != uuid.Nil {
		query = query.Where("root_id = ?", rootID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormCommentRepository) Patch(ctx context.Context, id uuid.UUID, comment *entities.Comment) error {
	tx := transaction.GetTx(ctx, r.db)

	result := tx.Model(&entities.Comment{}).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormCommentRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
