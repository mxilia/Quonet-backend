package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Handler     string    `gorm:"type:varchar(255);uniqueKey" json:"handler"` // Need to validate
	Email       string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	ProfileUrl  string    `gorm:"type:varchar(512);default:''" json:"profile_url"`
	Role        string    `gorm:"type:varchar(20);default:'member';check:role IN ('member','admin','owner')" json:"role"`
	IsBanned    bool      `gorm:"default:false" json:"is_banned"`
	BannedUntil time.Time `gorm:"timestamptz(3)" json:"banned_until"`

	Posts    []Post    `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"posts"`
	Comments []Comment `gorm:"foreignKey:AuthorID" json:"comments"`
	Likes    []Like    `gorm:"foreignKey:OwnerID" json:"likes"`

	CreatedAt time.Time `gorm:"timestamptz(3)" json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()

	}
	if u.Handler == "" {
		u.Handler = "user-" + uuid.NewString()[:8]
	}
	return nil
}
