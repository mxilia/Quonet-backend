package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/internal/transaction"
	"gorm.io/gorm"
)

type GormLikeRepository struct {
	db *gorm.DB
}

func NewGormLikeRepository(db *gorm.DB) LikeRepository {
	return &GormLikeRepository{db: db}
}

func (r *GormLikeRepository) Save(ctx context.Context, like *entities.Like) error {
	tx := transaction.GetTx(ctx, r.db)
	if err := tx.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormLikeRepository) Find(parentType string, ownerID uuid.UUID, parentID uuid.UUID, offset int, limit int) ([]*entities.Like, error) {
	var query *gorm.DB
	if parentType == "" {
		query = r.db
	} else {
		query = r.db.Where("parent_type = ?", parentType)
	}

	if ownerID != uuid.Nil {
		query = query.Where("owner_id = ?", ownerID)
	}

	if parentID != uuid.Nil {
		query = query.Where("parent_id = ?", parentID)
	}

	var likesValue []entities.Like
	if err := query.Limit(limit).Offset(offset).Find(&likesValue).Error; err != nil {
		return nil, err
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByID(id uuid.UUID) (*entities.Like, error) {
	var like entities.Like
	if err := r.db.First(&like, id).Error; err != nil {
		return nil, err
	}
	return &like, nil
}

func (r *GormLikeRepository) Count(parentType string, ownerID uuid.UUID, parentID uuid.UUID) (int64, error) {
	query := r.db.Model(&entities.Like{})
	if parentType != "" {
		query = r.db.Where("parent_type = ?", parentType)
	}

	if ownerID != uuid.Nil {
		query = query.Where("owner_id = ?", ownerID)
	}

	if parentID != uuid.Nil {
		query = query.Where("parent_id = ?", parentID)
	}

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormLikeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tx := transaction.GetTx(ctx, r.db)
	result := tx.Delete(&entities.Like{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
