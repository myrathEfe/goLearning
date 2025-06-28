package config

import (
	"fmt"
	"proje-api/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@postgresql_db:5432/youtube_student_api?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to open DB connection: %w", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get sql.DB: %w", err))
	}

	start := time.Now()
	timeout := 10 * time.Second

	for {
		err := sqlDB.Ping()
		if err == nil {
			break // bağlantı sağlandı
		}

		if time.Since(start) > timeout {
			panic(fmt.Errorf("failed to connect to DB after %v: %w", timeout, err))
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Connected to DB")
	DB = db

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Errorf("auto migrate failed: %w", err))
	}
}
