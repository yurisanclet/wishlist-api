package repositories

import (
	"backend/config"
	"backend/models"
	"errors"
)

type WishItemRepository interface {
	CreateWishItem(wishItem *models.WishItem) error
	GetWishItemsByWishlistId(wishlistId string) ([]models.WishItem, error)
	UpdateWishItem(id string, wishItem *models.WishItem) error
	DeleteWishItem(id string) error
}

type wishIteRepositoryImpl struct{}

func NewWishItemRepository() WishItemRepository {
	return &wishIteRepositoryImpl{}
}

func (w *wishIteRepositoryImpl) CreateWishItem(wishItem *models.WishItem) error {
	result := config.DB.Create(wishItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteWishItem implements WishItemRepository.
func (w *wishIteRepositoryImpl) DeleteWishItem(id string) error {
	result := config.DB.Where("id = ?", id).Delete(&models.WishItem{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("wishitem not found")
	}

	return nil
}

func (w *wishIteRepositoryImpl) GetWishItemsByWishlistId(wishlistId string) ([]models.WishItem, error) {
	var wishItems []models.WishItem
	result := config.DB.Where("wishlist_id = ?", wishlistId).Find(&wishItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return wishItems, nil
}

func (w *wishIteRepositoryImpl) UpdateWishItem(id string, wishItem *models.WishItem) error {
	updateData := map[string]interface{}{}

	if wishItem.Description != "" {
		updateData["description"] = wishItem.Description
	}

	if wishItem.Name != "" {
		updateData["name"] = wishItem.Name
	}

	updateData["is_bought"] = wishItem.IsBought

	if wishItem.Link != "" {
		updateData["link"] = wishItem.Link
	}

	if wishItem.Price != 0 {
		updateData["price"] = wishItem.Price
	}

	if wishItem.Priority != "" {
		updateData["priority"] = wishItem.Priority
	}

	if len(updateData) == 0 {
		return errors.New("no fields to update")
	}

	result := config.DB.Model(&models.WishItem{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("wish not found or no changes detected")
	}

	return nil
}