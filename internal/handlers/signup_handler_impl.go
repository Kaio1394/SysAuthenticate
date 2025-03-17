package handlers

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignupHandlerImpl struct {
	s *services.SignupServiceImpl
}

func NewSignupHandlerImpl(s *services.SignupServiceImpl) *SignupHandlerImpl {
	return &SignupHandlerImpl{s}
}

func (h *SignupHandlerImpl) RegisterNewUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.s.RegisterNewUser(context.Background(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
