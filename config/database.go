package config

import (
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func (c Config) ConnectDatabase(config Config) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Database connected")
}

// func getEnv(key, fallback string) string {
// 	if value, ok := os.LookupEnv(key); ok {
// 		print(value)
// 		return value
// 	}
// 	return fallback
// }
