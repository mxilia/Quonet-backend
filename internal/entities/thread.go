package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Thread struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"` // Need to validate
	Description string    `gorm:"text" json:"description"`
	ImageUrl    string    `gorm:"type:varchar(255)" json:"image_url"`

	Posts []Post `gorm:"foreignKey:ThreadID;constraint:OnDelete:CASCADE;" json:"posts"`

	CreatedAt time.Time `gorm:"type:timestamptz(3)" json:"created_at"`
}

func (u *Thread) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
