package seeders

import (
	"log"

	"backend/config"
	"backend/models"
)

func SeedItems() {
	var count int64
	config.DB.Model(&models.Item{}).Count(&count)

	if count > 0 {
		log.Println("Seeder: items already exist, skipping...")
		return
	}

	items := []models.Item{
		{
			Name:  "Laptop",
			Stock: 10,
			Price: 15000000,
		},
		{
			Name:  "Mouse",
			Stock: 50,
			Price: 150000,
		},
		{
			Name:  "Keyboard",
			Stock: 30,
			Price: 300000,
		},
	}

	if err := config.DB.Create(&items).Error; err != nil {
		log.Println("Seeder items error:", err)
		return
	}

	log.Println("Seeder: items created")
}
