package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
	"gorm.io/gorm"
)

type GormAnnouncementRepository struct {
	db *gorm.DB
}

func NewGormAnnouncementRepository(db *gorm.DB) AnnouncementRepository {
	return &GormAnnouncementRepository{db: db}
}

func (r *GormAnnouncementRepository) Save(announcement *entities.Announcement) error {
	if err := r.db.Save(announcement).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormAnnouncementRepository) Find(offset int, limit int) ([]*entities.Announcement, error) {
	var announcementsValue []entities.Announcement
	if err := r.db.Limit(limit).Offset(offset).Preload("Author").Order("created_at DESC").Find(&announcementsValue).Error; err != nil {
		return nil, err
	}

	announcements := make([]*entities.Announcement, len(announcementsValue))
	for i := range announcementsValue {
		announcements[i] = &announcementsValue[i]
	}
	return announcements, nil
}

func (r *GormAnnouncementRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&entities.Announcement{}).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

func (r *GormAnnouncementRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&entities.Announcement{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
