package database

import (
	"context"
	"fmt"
	"log"
	"time"
	u "user/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	NewDatabaseConnection() (*gorm.DB, error)
}

func NewDatabaseConnection(automigrate bool) (*gorm.DB, error) {
	// godotenv package to get .env variables.
	envFile, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s \n",
		envFile["databaseHost"],
		envFile["databaseUser"],
		envFile["databasePassword"],
		envFile["databaseName"],
		envFile["databasePort"],
		envFile["databaseSSLMode"],
	)
	fmt.Println("dsn: ",dsn)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err: ",err)
		panic("Database connection failed")
	}

	if automigrate {
		database.AutoMigrate(&u.User{})
		if err != nil {
			fmt.Println("err: ",err)
			panic("Database failed to automigrate users")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()		
	
	dbInstance, _ := database.DB()
    err = dbInstance.PingContext(ctx)
    if err != nil {
        return nil, err
    }

	fmt.Println("Successfully connected to the database")

	return database, nil
}