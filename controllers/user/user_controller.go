package userController

import (
	dtos "backend/models/dtos/request"
	"backend/repositories"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userRepository = repositories.NewUserRepository()
var userService = services.NewUserService(userRepository)


func RegisterHandler(c *gin.Context) {
	var userDto dtos.UserRegisterDTO

	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	user, err := userService.RegisterUser(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário cadastrado com sucesso",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}