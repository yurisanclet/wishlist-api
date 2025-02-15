package services

import (
	"backend/models"
	dtos "backend/models/dtos/request"
	"backend/repositories"
	"fmt"
	"time"
)

type WishListService interface {
	CreateWishlist(userEmail string, dto dtos.WishlistRequestDto) (*models.Wishlist, error)
	GetWishlistsById(id string) (*models.Wishlist, error)
	GetWishlistsByUserId(userEmail string) ([]models.Wishlist, error)
	UpdateWishlist(id string, dto dtos.WishlistRequestDto) error
	DeleteWishlist(id string) error
}

type wishlistServiceImpl struct {
	wishlistRepository repositories.WishListRepository
	userRepository repositories.UserRepository
}

func NewWishlistService(
	wishlistRepo repositories.WishListRepository,
	userRepository repositories.UserRepository) WishListService {
	return &wishlistServiceImpl{
		wishlistRepository: wishlistRepo,
		userRepository: userRepository,
	}
}

func (w *wishlistServiceImpl) CreateWishlist(userEmail string, dto dtos.WishlistRequestDto) (*models.Wishlist, error) {
	existingUser, err := w.userRepository.GetUserByEmail(userEmail)
	// nil representa ausencia de valor para tipos de referencia
	// existingUser != nil significa que user foi encontrado
	if err != nil {
		return nil, err
	}

	eventDate, err := time.Parse("2006-01-02",dto.EventDate)
	if err != nil {
		return nil, err
	}

	newWishlist := &models.Wishlist{
		UserID: existingUser.ID,
		Title: dto.Title,
		Description: dto.Description,
		EventDate: eventDate,
	}
	fmt.Print(newWishlist)

	err = w.wishlistRepository.CreateWishList(newWishlist)
	if err != nil {
		return nil, err
	}

	return newWishlist, nil
}

func (w *wishlistServiceImpl) DeleteWishlist(id string) error {
	return w.wishlistRepository.DeleteWishList(id)
}

func (w *wishlistServiceImpl) GetWishlistsById(id string) (*models.Wishlist, error) {
	return w.wishlistRepository.GetWishListByid(id)
}

func (w *wishlistServiceImpl) GetWishlistsByUserId(userEmail string) ([]models.Wishlist, error) {
	existingUser, err := w.userRepository.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	wishlists, err := w.wishlistRepository.GetWishListsByUserId(existingUser.ID)
	if err != nil {
		return nil, err
	}

	return wishlists, nil
}

func (w *wishlistServiceImpl) UpdateWishlist(id string, dto dtos.WishlistRequestDto) error {

	eventDate, err := time.Parse("2006-01-02",dto.EventDate)
	if err != nil {
		return err
	}

	return w.wishlistRepository.UpdateWishList(id, &models.Wishlist{
		Title: dto.Title,
		Description: dto.Description,
		EventDate: eventDate,
	})
}


