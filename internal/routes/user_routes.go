package routes

import (
	"SysAuthenticate/internal/handlers"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(c *gin.Engine, db *gorm.DB) {
	r := repository.NewUserRepositoryImpl(db)
	s := services.NewUserServiceImpl(r)
	h := handlers.NewUserHandlerImpl(s)
	c.GET("/users", h.GetUsers)
}
