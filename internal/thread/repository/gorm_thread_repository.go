package repository

import (
	"github.com/mxilia/Conflux-backend/internal/entities"
	"gorm.io/gorm"
)

type GormThreadRepository struct {
	db *gorm.DB
}

func NewGormThreadRepository(db *gorm.DB) ThreadRepository {
	return &GormThreadRepository{db: db}
}

func (r *GormThreadRepository) Save(thread *entities.Thread) error {
	if err := r.db.Create(thread).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormThreadRepository) FindAll() ([]*entities.Thread, error) {
	var threadsValue []entities.Thread
	if err := r.db.Find(&threadsValue).Error; err != nil {
		return nil, err
	}

	threads := make([]*entities.Thread, len(threadsValue))
	for i := range threadsValue {
		threads[i] = &threadsValue[i]
	}
	return threads, nil
}

func (r *GormThreadRepository) FindByID(id uint) (*entities.Thread, error) {
	var thread entities.Thread
	if err := r.db.First(&thread, id).Error; err != nil {
		return nil, err
	}
	return &thread, nil
}

func (r *GormThreadRepository) Delete(id uint) error {
	result := r.db.Delete(&entities.Thread{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
