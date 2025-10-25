package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	Email      string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	ProfileUrl string    `gorm:"type:varchar(512)" json:"profile_url"`
	IsAdmin    bool      `gorm:"default:false" json:"is_admin"`

	Posts    []Post    `gorm:"foreignKey:AuthorID" json:"posts"`
	Comments []Comment `gorm:"foreignKey:AuthorID" json:"comments"`

	CreatedAt time.Time `gorm:"timestamp" json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
