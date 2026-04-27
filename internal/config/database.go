package config

import (
	"fmt"
	"log"
	"os"

	"reast-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Auto Migrate the models
	err = db.AutoMigrate(&models.UserCredential{}, &models.UserProfile{})
	if err != nil {
		log.Fatal("Failed to migrate database!", err)
	}

	fmt.Println("Database connection & migration successful")
	return db
}
