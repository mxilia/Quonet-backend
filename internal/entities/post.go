package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;index:index_created_id;priority:2" json:"id"`
	Title        string    `gorm:"type:varchar(255)" json:"title"`
	AuthorID     uuid.UUID `gorm:"type:uuid" json:"author_id"`
	ThreadID     uuid.UUID `gorm:"type:uuid" json:"thread_id"`
	Content      string    `gorm:"type:text" json:"content"`
	ThumbnailUrl string    `gorm:"type:varchar(512)" json:"thumbnail_url"`

	Like    uint `gorm:"default:0" json:"like"`
	Dislike uint `gorm:"default:0" json:"dislike"`

	Comments []Comment `gorm:"foreignKey:RootID" json:"comments"`

	CreatedAt time.Time `gorm:"type:timestamp;index:index_created_id;priority:1" json:"created_at"`
}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
