package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
	"gorm.io/gorm"
)

type GormLikeRepository struct {
	db *gorm.DB
}

func NewGormLikeRepository(db *gorm.DB) LikeRepository {
	return &GormLikeRepository{db: db}
}

func (r *GormLikeRepository) Save(like *entities.Like) error {
	if err := r.db.Create(like).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormLikeRepository) FindAll() ([]*entities.Like, error) {
	var likesValue []entities.Like
	if err := r.db.Find(&likesValue).Error; err != nil {
		return nil, err
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByOwnerID(id uuid.UUID) ([]*entities.Like, error) {
	var likesValue []entities.Like
	if err := r.db.Where("owner_id = ?", id).Find(&likesValue).Error; err != nil {
		return nil, err
	}

	likes := make([]*entities.Like, len(likesValue))
	for i := range likesValue {
		likes[i] = &likesValue[i]
	}
	return likes, nil
}

func (r *GormLikeRepository) FindByParentID(id uuid.UUID) ([]*entities.Like, error) {
	var likesValue []entities.Like
	if err := r.db.Where("parent_id = ?", id).Find(&likesValue).Error; err != nil {
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

func (r *GormLikeRepository) CountByParentID(parentType string, id uuid.UUID) (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Like{}).Where("parent_type = ? and parent_id = ?", parentType, id).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormLikeRepository) IsParentLikedByMe(parentType string, parentID uuid.UUID, myID uuid.UUID) (bool, error) {
	var count int64
	if err := r.db.Model(&entities.Like{}).Where("parent_type = ? AND parent_id = ? AND owner_id = ?", parentType, parentID, myID).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *GormLikeRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Like{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
