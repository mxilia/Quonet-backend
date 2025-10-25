package entities

import (
	"time"
)

type Thread struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `gorm:"type:varchar(255);unique" json:"title"`

	Posts []Post `gorm:"foreignKey:ThreadID" json:"posts"`

	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
}
