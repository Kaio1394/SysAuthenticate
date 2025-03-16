package utils

import "SysAuthenticate/internal/models"

func ValidateToken(token string) bool {
	return true
}

func GenerateToken(user models.User) string {
	return ""
}
