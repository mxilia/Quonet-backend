package repository

import (
	"context"
	"errors"

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

func (r *GormLikeRepository) FindAll(parentType string) ([]*entities.Like, error) {
	var likesValue []entities.Like
	if parentType == "" {
		if err := r.db.Find(&likesValue).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("parent_type = ?", parentType).Find(&likesValue).Error; err != nil {
			return nil, err
		}
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByOwnerID(parentType string, id uuid.UUID) ([]*entities.Like, error) {
	var likesValue []entities.Like
	if parentType == "" {
		if err := r.db.Where("owner_id = ?", id).Find(&likesValue).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("parent_type = ? and owner_id = ?", parentType, id).Find(&likesValue).Error; err != nil {
			return nil, err
		}
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByParentID(parentType string, id uuid.UUID) ([]*entities.Like, error) {
	var likesValue []entities.Like
	if parentType == "" {
		if err := r.db.Where("parent_id = ?", id).Find(&likesValue).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("parent_type = ? and parent_id = ?", parentType, id).Find(&likesValue).Error; err != nil {
			return nil, err
		}
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByParentIDAndOwnerID(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (*entities.Like, error) {
	var like entities.Like
	if err := r.db.Where("parent_type = ? and parent_id = ? and owner_id = ?", parentType, parentID, ownerID).First(&like).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &like, nil
}

func (r *GormLikeRepository) FindByID(parentType string, id uuid.UUID) (*entities.Like, error) {
	var like entities.Like
	if parentType == "" {
		if err := r.db.First(&like, id).Error; err != nil {
			return nil, err
		}
	} else {
		if err := r.db.Where("parent_type = ?", parentType).First(&like, id).Error; err != nil {
			return nil, err
		}
	}
	return &like, nil
}

func (r *GormLikeRepository) CountByParentID(parentType string, id uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Like{}).Where("parent_type = ? and parent_id = ?", parentType, id).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormLikeRepository) IsParentLikedByMe(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Like{}).Where("parent_type = ? AND parent_id = ? AND owner_id = ?", parentType, parentID, ownerID).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *GormLikeRepository) Delete(ctx context.Context, parentType string, id uuid.UUID) error {
	tx := transaction.GetTx(ctx, r.db)

	var result *gorm.DB
	if parentType == "" {
		result = tx.Delete(&entities.Like{}, id)
	} else {
		result = tx.Where("parent_type = ?", parentType).Delete(&entities.Like{}, id)
	}

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
