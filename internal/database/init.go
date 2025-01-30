package database

import (
	"log"
	"os"

	"github.com/abdulridhoramadhan/CMS-Go-Project/cms-server/internal/models"
	"gorm.io/gorm"
)

func InitDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	var adminCount int64
	db.Model(&models.User{}).Where("email = ? AND role = ?", "admin@mail.com", "Admin").Count(&adminCount)

	if adminCount == 0 {
		adminPassword := os.Getenv("ADMIN_DEFAULT_PASSWORD")
		log.Printf("Creating admin user with password from env: %s", adminPassword)
		if adminPassword == "" {
			log.Fatal("ADMIN_DEFAULT_PASSWORD environment variable is required")
		}

		admin := models.User{
			Username: "admin",
			Email:    "admin@mail.com",
			Password: adminPassword,
			Role:     "Admin",
		}

		if err := db.Create(&admin).Error; err != nil {
			log.Fatal("Failed to create admin user:", err)
		}
		log.Println("Default admin user created")
	}
} 