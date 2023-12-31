package models

import (
	"time"
)

type Photo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `gorm:"not null" json:"photoUrl"`
	UserID    uint      `gorm:"not null" json:"userId"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
}
