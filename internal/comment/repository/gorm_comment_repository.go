package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
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

func (r *GormCommentRepository) FindAll() ([]*entities.Comment, error) {
	var commentsValue []entities.Comment
	if err := r.db.Find(&commentsValue).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentsValue))
	for i := range commentsValue {
		comments[i] = &commentsValue[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByAuthorID(id uuid.UUID) ([]*entities.Comment, error) {
	var commentsValue []entities.Comment
	if err := r.db.Where("author_id = ?", id).Find(&commentsValue).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentsValue))
	for i := range commentsValue {
		comments[i] = &commentsValue[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByParentID(id uuid.UUID) ([]*entities.Comment, error) {
	var commentsValue []entities.Comment
	if err := r.db.Where("parent_id = ?", id).Find(&commentsValue).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentsValue))
	for i := range commentsValue {
		comments[i] = &commentsValue[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByRootID(id uuid.UUID) ([]*entities.Comment, error) {
	var commentsValue []entities.Comment
	if err := r.db.Where("root_id = ?", id).Find(&commentsValue).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentsValue))
	for i := range commentsValue {
		comments[i] = &commentsValue[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByID(id uuid.UUID) (*entities.Comment, error) {
	var comment entities.Comment
	if err := r.db.Where("id = ?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *GormCommentRepository) Patch(id uuid.UUID, comment *entities.Comment) error {
	result := r.db.Model(&entities.Comment{}).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
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
