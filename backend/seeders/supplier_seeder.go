package seeders

import (
	"log"

	"backend/config"
	"backend/models"
)

func SeedSuppliers() {
	var count int64
	config.DB.Model(&models.Supplier{}).Count(&count)

	if count > 0 {
		log.Println("Seeder: suppliers already exist, skipping...")
		return
	}

	suppliers := []models.Supplier{
		{
			Name:    "PT Sumber Makmur",
			Email:   "supplier1@mail.com",
			Address: "Jakarta",
		},
		{
			Name:    "CV Jaya Abadi",
			Email:   "supplier2@mail.com",
			Address: "Bandung",
		},
	}

	if err := config.DB.Create(&suppliers).Error; err != nil {
		log.Println("Seeder suppliers error:", err)
		return
	}

	log.Println("Seeder: suppliers created")
}
