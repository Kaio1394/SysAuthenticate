package routes

import (
	"SysAuthenticate/internal/handlers"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterLoginRoutes(r *gin.Engine, db *gorm.DB) {
	repoRole := repository.NewRoleRepositoryImpl(db)
	repoUserRole := repository.NewUserRoleRepositoryImpl(db)
	repoSignup := repository.NewSignupRepositoryImpl(db)

	serviceLogin := services.NewLoginServiceImpl(repoSignup, repoUserRole, repoRole)
	h := handlers.NewLoginHandlerImpl(serviceLogin)

	r.POST("/login", h.Login)
}
