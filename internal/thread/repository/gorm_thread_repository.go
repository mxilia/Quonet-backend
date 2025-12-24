package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
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

func (r *GormThreadRepository) Find(title string, offset int, limit int) ([]*entities.Thread, error) {
	query := r.db.Limit(limit).Offset(offset)

	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	var threadsValue []entities.Thread
	if err := query.Find(&threadsValue).Error; err != nil {
		return nil, err
	}

	threads := make([]*entities.Thread, len(threadsValue))
	for i := range threadsValue {
		threads[i] = &threadsValue[i]
	}
	return threads, nil
}

func (r *GormThreadRepository) FindByID(id uuid.UUID) (*entities.Thread, error) {
	var thread entities.Thread
	if err := r.db.First(&thread, id).Error; err != nil {
		return nil, err
	}
	return &thread, nil
}

func (r *GormThreadRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Thread{}).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormThreadRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Thread{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
