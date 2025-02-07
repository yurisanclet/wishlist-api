package services

import (
	"backend/models"
	dtos "backend/models/dtos/request"
	"backend/repositories"
	u "backend/utils"
	"errors"
)

type AuthService interface {
	AuthenticateUser(dto dtos.LoginRequest) (*models.User, error)
}

type authServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (a *authServiceImpl) AuthenticateUser(dto dtos.LoginRequest) (*models.User, error) {
	user, err := a.userRepo.GetUserByEmail(dto.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !u.CheckPassword(user.Password, dto.Password) {
		return nil, errors.New("credenciais inválidas")
	}

	return user, nil
}


