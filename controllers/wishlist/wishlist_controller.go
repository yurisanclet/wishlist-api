package wishlistController

import (
	dtos "backend/models/dtos/request"
	"backend/repositories"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var wishListRepository = repositories.NewWishListRepository()
var userRepository = repositories.NewUserRepository()
var wishListService = services.NewWishlistService(wishListRepository, userRepository)


func CreateWishList(c *gin.Context) {
	var wishListDto dtos.WishlistRequestDto
	wishListDto.IsUpdate = false 

	userEmail, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
	}

	if err := c.ShouldBindJSON(&wishListDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wishList, err := wishListService.CreateWishlist(userEmail.(string), wishListDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, wishList)
}


func GetWishListById(c *gin.Context) {
	
}


