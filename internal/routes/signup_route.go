package routes

import (
	"SysAuthenticate/internal/handlers"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterSignupRoutes(c *gin.Engine, db *gorm.DB) {
	r := repository.NewSignupRepositoryImpl(db)
	s := services.NewSignupServiceImpl(r)
	h := handlers.NewSignupHandler(s)
	c.POST("/signup", h.RegisterNewUser)
}
