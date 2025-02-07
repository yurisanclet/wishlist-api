package auth

import (
	dtos "backend/models/dtos/request"
	"backend/repositories"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userRepo = repositories.NewUserRepository()
var authService = services.NewAuthService(userRepo)
var jwtService = services.NewJWTService()


func LoginHandler(c *gin.Context){
	var LoginReq dtos.LoginRequest

	if err := c.ShouldBindJSON(&LoginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	 
	user, err := authService.AuthenticateUser(LoginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := jwtService.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}