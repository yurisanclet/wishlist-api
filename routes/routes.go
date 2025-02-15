package routes

import (
	"backend/controllers/auth"
	userController "backend/controllers/user"
	wishItemController "backend/controllers/wishitem"
	wishlistController "backend/controllers/wishlist"
	"backend/middlewares"
	"backend/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	jwtService := services.NewJWTService()

	userRoutes := router.Group("/api/user")
	{
		userRoutes.POST("/register", userController.RegisterHandler)
	}

	wishListRoutes := router.Group("/api/wishlist")
	{
		wishListRoutes.POST("/create", middlewares.AuthMiddleware(jwtService), wishlistController.CreateWishList)
	}

	wishItemRoutes := router.Group("/api/wishitem")
	{
		wishItemRoutes.POST("/create", middlewares.AuthMiddleware(jwtService), wishItemController.CreateWishItem)
	}

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", auth.LoginHandler)
	}
}