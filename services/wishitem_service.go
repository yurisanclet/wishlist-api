package services

import (
	"backend/models"
	dtos "backend/models/dtos/request"
	"backend/repositories"
)

type WishItemService interface {
	CreateWishItem(dto dtos.WishItemRequestDto) (*models.WishItem, error)
	GetWishItemsByWishlistId(wishlistId string) ([]models.WishItem, error)
	UpdateWishItem(id string, dto dtos.WishItemRequestDto) error
	DeleteWishItem(id string) error
}

type wishItemServiceImpl struct {
	wishlistRepository repositories.WishListRepository
	wishItemRepository repositories.WishItemRepository
}

func NewWishItemService(
	wishlistRepository repositories.WishListRepository,
	wishItemRepository repositories.WishItemRepository) WishItemService {
	return &wishItemServiceImpl{
		wishlistRepository: wishlistRepository,
		wishItemRepository: wishItemRepository}
}

func (w *wishItemServiceImpl) CreateWishItem(dto dtos.WishItemRequestDto) (*models.WishItem, error) {
	_, err := w.wishlistRepository.GetWishListByid(dto.WishListID)
	if err != nil {
		return nil, err
	}

	newWishitem := &models.WishItem{
		WishlistID: dto.WishListID,
		Name: dto.Name,
		Description: dto.Description,
		Priority: dto.Priority,
		Price: dto.Price,
		Link: dto.Link,
		IsBought: dto.IsBought,
	}

	err = w.wishItemRepository.CreateWishItem(newWishitem)
	
	if err != nil {
		return nil, err
	}

	return newWishitem, nil
}

func (w *wishItemServiceImpl) DeleteWishItem(id string) error {
	return w.wishItemRepository.DeleteWishItem(id)
}

func (w *wishItemServiceImpl) GetWishItemsByWishlistId(wishlistId string) ([]models.WishItem, error) {
	existingWishlist, err := w.wishlistRepository.GetWishListByid(wishlistId)
	if err != nil {
		return nil, err
	}

	wishItems, err := w.wishItemRepository.GetWishItemsByWishlistId(existingWishlist.ID)
	if err != nil {
		return nil, err
	}

	return wishItems, nil
}

func (w *wishItemServiceImpl) UpdateWishItem(id string, dto dtos.WishItemRequestDto) error {
	return w.wishItemRepository.UpdateWishItem(id, &models.WishItem{
		WishlistID: dto.WishListID,
		Name: dto.Name,
		Description: dto.Description,
		Priority: dto.Priority,
		Price: dto.Price,
		Link: dto.Link,
		IsBought: dto.IsBought,
	})
}


