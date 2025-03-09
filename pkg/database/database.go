package database

import (
	"fmt"
	"os"

	"dribbble-clone-be/internal/auth"
	"dribbble-clone-be/internal/profile"
	"dribbble-clone-be/internal/shot"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the schemas
	err = db.AutoMigrate(&auth.User{}, &profile.Profile{}, &shot.Shot{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
