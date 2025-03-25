package routes

import (
	"SysAuthenticate/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	h := handlers.NewAuthHandlerImpl()

	r.POST("/validate/token", h.ValidateToken())
}
