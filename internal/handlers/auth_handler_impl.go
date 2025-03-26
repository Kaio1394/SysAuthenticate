package handlers

import (
	"SysAuthenticate/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthHandlerImpl struct{}

func NewAuthHandlerImpl() *AuthHandlerImpl {
	return &AuthHandlerImpl{}
}

func (h *AuthHandlerImpl) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{"valid": true, "email": claims.Email})
	}
}
