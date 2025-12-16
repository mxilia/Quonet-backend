package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;index:index_created_id;priority:2" json:"id"`
	Title        string    `gorm:"type:varchar(255)" json:"title"` // Need to validate
	AuthorID     uuid.UUID `gorm:"type:uuid" json:"author_id"`
	ThreadID     uuid.UUID `gorm:"type:uuid" json:"thread_id"`
	Content      string    `gorm:"type:text" json:"content"`               // Need to validate
	ThumbnailUrl string    `gorm:"type:varchar(512)" json:"thumbnail_url"` // Need to validate
	IsPrivate    bool      `gorm:"default:false" json:"is_private"`
	LikeCount    int64     `gorm:"default:0" json:"like_count"`

	Author   User      `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"author"`
	Thread   Thread    `gorm:"foreignKey:ThreadID;constraint:OnDelete:CASCADE" json:"thread"`
	Likes    []Like    `gorm:"polymorphic:Parent;polymorphicValue:post;constraint:OnDelete:CASCADE" json:"likes"`
	Comments []Comment `gorm:"foreignKey:RootID;constraint:OnDelete:CASCADE" json:"comments"`

	CreatedAt time.Time `gorm:"type:timestamptz(3);index:index_created_id;priority:1" json:"created_at"`
}

func (u *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
