package routes

import (
	"backend/controllers/auth"
	userController "backend/controllers/user"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	jwtService := services.NewJWTService()

	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("/register", middlewares.AuthMiddleware(jwtService) ,userController.RegisterHandler)
	}

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", auth.LoginHandler)
	}
}