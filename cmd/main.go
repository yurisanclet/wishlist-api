package main

import (
	"backend/config"
	"backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	r := gin.Default()

	err := godotenv.Load("D:/IWish/wishlist-api/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
}

	config := config.Config{
		Host:     os.Getenv("DB_HOST"),
		User:		 os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	config.ConnectDatabase(config)

	routes.RegisterRoutes(r)

	port := ":8080"
	fmt.Println("Servidor rodando na porta", port)

	r.Run()
}