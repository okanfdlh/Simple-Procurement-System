package seeders

import (
	"log"

	"backend/config"
	"backend/models"
	"backend/utils"
)

func SeedUsers() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("Seeder: users already exist, skipping...")
		return
	}

	users := []models.User{
		{
			Username: "admin",
			Password: utils.HashPassword("password123"),
			Role:     "admin",
		},
		{
			Username: "staff",
			Password: utils.HashPassword("password123"),
			Role:     "staff",
		},
	}

	if err := config.DB.Create(&users).Error; err != nil {
		log.Println("Seeder users error:", err)
		return
	}

	log.Println("Seeder: users created")
}
