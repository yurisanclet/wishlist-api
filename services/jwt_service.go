package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string) (string, error)
	ValidateToken(token string) (string, error)
}

type jwtServiceImpl struct{
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	secret := os.Getenv("JWT_SECRET")
	return &jwtServiceImpl{
		secretKey: secret,
		issuer: "wishlist-api", 
	}
}

func (j *jwtServiceImpl) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"iss": j.issuer,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))   
}

func (j *jwtServiceImpl) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("❌ ERRO: Método de assinatura inválido")
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("❌ ERRO AO VALIDAR TOKEN:", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New("email inválido no token")
		}
		return email, nil
	}

	return "", errors.New("token inválido")
}
