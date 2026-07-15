package routes
import (
	"github.com/AbhayBajpai10/url-shortener/internal/url/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *handler.Handler) {
	r.POST("/",h.CreateURL)
	r.GET("/", h.GetMyURLs)
	r.DELETE("/:id", h.DeleteURL)
}