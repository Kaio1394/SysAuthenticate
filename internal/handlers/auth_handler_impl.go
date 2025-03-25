package handlers

import (
	"SysAuthenticate/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlerImpl struct{}

func NewAuthHandlerImpl() *AuthHandlerImpl {
	return &AuthHandlerImpl{}
}

func (h *AuthHandlerImpl) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("token")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{"valid": true, "email": claims.Email})
	}
}
