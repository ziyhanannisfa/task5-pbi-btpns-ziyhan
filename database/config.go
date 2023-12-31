package database

import (
	"PBI/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "user=postgres dbname=photo password=zi00911 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = db
}

func Migrate() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
	fmt.Println("Database migrated")
}
