package utils

import (
	"SysAuthenticate/internal/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("seu_segredo_super_secreto")

func ValidateJWT(tokenString string) (*models.UserClaims, error) {
	claims := &models.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}

func GenerateToken(user *models.User, role string, originType string) (string, error) {
	userClaims := models.UserClaims{
		UserID: user.Id,
		Email:  user.Email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    originType,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
