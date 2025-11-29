package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OwnerID    uuid.UUID `gorm:"type:uuid" json:"owner_id"`
	ParentID   uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	ParentType string    `gorm:"type:VARCHAR(255)" json:"parent_type"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
}

func (u *Like) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
