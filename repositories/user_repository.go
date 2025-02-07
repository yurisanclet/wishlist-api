package repositories

import (
	"backend/config"
	"backend/models"
	"errors"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
}

type userRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func (u *userRepositoryImpl) CreateUser(user *models.User) error {
	result := config.DB.Create(user)
	if result != nil {
		return result.Error
	}
	return nil
}

func (u *userRepositoryImpl) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}

func (u *userRepositoryImpl) GetUserById(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return nil, errors.New("usuário não encontrado")
	}
	return &user, nil
}

