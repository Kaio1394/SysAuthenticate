package routes

import (
	"SysAuthenticate/internal/handlers"
	"SysAuthenticate/internal/repository"
	"SysAuthenticate/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoleRoutes(r *gin.Engine, db *gorm.DB) {
	rr := repository.NewRoleRepositoryImpl(db)
	ru := repository.NewUserRepositoryImpl(db)

	repo := repository.NewUserRoleRepositoryImpl(db)
	serv := services.NewUserRoleServiceImpl(repo, rr, ru)
	h := handlers.NewUserRoleHandlerImpl(serv)

	r.POST("/user/role", h.CreateNewUserRole)
}
