package main

import (
	"backend/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	port := ":8080"
	fmt.Println("Servidor rodando na porta", port)

	r.Run()
}