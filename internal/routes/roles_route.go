package routes

import (
	"SysAuthenticate/internal/handlers"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoleRoutes(r *gin.Engine, db *gorm.DB) {
	repo := repository.NewRoleRepositoryImpl(db)
	s := services.NewRoleServiceImpl(repo)
	h := handlers.NewRoleHandlerImpl(s)

	r.POST("/role", h.CreateNewRole)
}
