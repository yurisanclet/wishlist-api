package wishItemController

import (
	dtos "backend/models/dtos/request"
	"backend/repositories"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)


var wishItemRepository = repositories.NewWishItemRepository()
var wishListRepository = repositories.NewWishListRepository()
var wishItemService = services.NewWishItemService(wishListRepository, wishItemRepository)


func CreateWishItem(c *gin.Context){
	var wishItemDto dtos.WishItemRequestDto
	wishItemDto.IsUpdate = false

	if err := c.ShouldBindJSON(&wishItemDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wishItem, err := wishItemService.CreateWishItem(wishItemDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wishItem) 
}