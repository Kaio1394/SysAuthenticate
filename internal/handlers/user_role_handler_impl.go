package handlers

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRoleHandlerImpl struct {
	s *services.UserRoleServiceImpl
}

func NewUserRoleHandlerImpl(s *services.UserRoleServiceImpl) *UserRoleHandlerImpl {
	return &UserRoleHandlerImpl{s}
}

func (h *UserRoleHandlerImpl) CreateNewUserRole(c *gin.Context) {
	var userRole models.UserRole
	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.s.CreateUserRole(context.Background(), &userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User role created"})
}
