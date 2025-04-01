package handlers

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandlerImpl struct {
	s *services.LoginServiceImpl
}

func NewLoginHandlerImpl(s *services.LoginServiceImpl) *LoginHandlerImpl {
	return &LoginHandlerImpl{s}
}

func (h *LoginHandlerImpl) Login(c *gin.Context) {
	var userLogin models.UserLogin
	originType := c.Request.Header.Get("Origin")
	if originType == "" || originType == "null" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Header 'origin' is empty!"})
		return
	}
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.s.Signup(context.Background(), &userLogin, originType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jwt_token":  token,
		"username":   userLogin.Username,
		"expires_in": 86400,
	})
}
