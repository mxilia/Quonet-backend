package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;index:index_created_id;priority:4" json:"id"`
	AuthorID uuid.UUID `gorm:"type:uuid" json:"author_id"`
	Content  string    `gorm:"type:text" json:"content"`

	ParentID uuid.UUID `gorm:"type:uuid;index:index_created_id;priority:2" json:"parent_id"`
	RootID   uuid.UUID `gorm:"type:uuid;index:index_created_id;priority:1" json:"root_id"`

	Author   User      `gorm:"foreignKey:AuthorID" json:"author"`
	Likes    []Like    `gorm:"polymorphic:Parent;polymorphicValue:post;constraint:OnDelete:CASCADE" json:"likes"`
	Comments []Comment `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE;" json:"comments"`

	CreatedAt time.Time `gorm:"type:timestamp;index:index_created_id;priority:3" json:"created_at"`
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
