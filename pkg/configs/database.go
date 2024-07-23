package configs

import (
	"donation-backend/pkg/helpers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	dbHost := helpers.GetEnv("DB_HOST", "localhost")
	dbPort := helpers.GetEnv("DB_PORT", "5432")
	dbName := helpers.GetEnv("DB_NAME", "secret_chat")
	dbUser := helpers.GetEnv("DB_USER", "root")
	dbPass := helpers.GetEnv("DB_PASS", "root")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}
