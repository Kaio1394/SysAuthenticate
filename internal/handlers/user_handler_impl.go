package handlers

import (
	"SysAuthenticate/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlerImpl struct {
	s *services.UserServiceImpl
}

func NewUserHandlerImpl(s *services.UserServiceImpl) *UserHandlerImpl {
	return &UserHandlerImpl{s}
}

func (h *UserHandlerImpl) GetUsers(c *gin.Context) {
	users, err := h.s.GetUsers(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
