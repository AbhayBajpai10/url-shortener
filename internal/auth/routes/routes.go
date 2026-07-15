package routes

import (
	"fmt"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.RouterGroup, h *handler.Handler) {
	fmt.Println("Registering auth routes")
	r.POST("/register", h.Register)
	r.POST("/login",h.Login)
}
