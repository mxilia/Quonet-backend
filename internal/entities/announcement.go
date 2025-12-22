package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Announcement struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;index:index_created_id;priority:2" json:"id"`
	AuthorID  uuid.UUID `gorm:"type:uuid" json:"author_id"`
	Content   string    `gorm:"type:text" json:"content"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
	CreatedAt time.Time `gorm:"type:timestamptz(3);index:index_created_id;priority:1" json:"created_at"`
}

func (u *Announcement) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
