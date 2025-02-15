package services

import (
	"backend/models"
	dtos "backend/models/dtos/request"
	"backend/repositories"
	u "backend/utils"
	"errors"
)

type UserService interface {
	RegisterUser(user *dtos.UserRegisterDTO) (*models.User, error)
}

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}

func (s *userServiceImpl) RegisterUser(user *dtos.UserRegisterDTO) (*models.User, error) {
	existingUser, _ := s.userRepository.GetUserByEmail(user.Email)
	if existingUser != nil {
		return nil, errors.New("usuário já cadastrado")
	}

	hashedPassword, err := u.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = s.userRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
