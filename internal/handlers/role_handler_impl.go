package handlers

import (
	"SysAuthenticate/internal/models"
	"SysAuthenticate/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandlerImpl struct {
	s *services.RoleServiceImpl
}

func NewRoleHandlerImpl(s *services.RoleServiceImpl) *RoleHandlerImpl {
	return &RoleHandlerImpl{s}
}

func (h *RoleHandlerImpl) CreateNewRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.s.CreateNewRole(context.Background(), &role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
