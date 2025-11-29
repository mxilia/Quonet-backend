package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Thread struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title string    `gorm:"type:varchar(255);unique" json:"title"` // Need to validate

	Posts []Post `gorm:"foreignKey:ThreadID" json:"posts"`

	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}

func (u *Thread) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
