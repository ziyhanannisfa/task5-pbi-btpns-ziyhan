package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	Photos    []Photo   `gorm:"foreignKey:UserID" json:"photos"`
}

// LoginData represents the data required for user login
type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Claims represents the claims in the JWT token
type Claims struct {
	UserID uint `json:"userId"`
	jwt.StandardClaims
}

// BeforeSave hooks into the GORM lifecycle and is triggered before saving to the database
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Implement any logic before saving the user to the database
	return nil
}
