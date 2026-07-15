package router

import (
	"github.com/AbhayBajpai10/url-shortener/config"
	authHandler "github.com/AbhayBajpai10/url-shortener/internal/auth/handler"
	authRoutes "github.com/AbhayBajpai10/url-shortener/internal/auth/routes"
	middleware "github.com/AbhayBajpai10/url-shortener/internal/middleware"
	urlHandler "github.com/AbhayBajpai10/url-shortener/internal/url/handler"
	urlRoutes "github.com/AbhayBajpai10/url-shortener/internal/url/routes"
	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *authHandler.Handler, urlHandler *urlHandler.Handler, cfg *config.Config) *gin.Engine {

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
	r.GET("/:shortCode", urlHandler.RedirectURL)
	api := r.Group("/api/v1")
	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	auth := api.Group("/auth")

	authRoutes.RegisterRoute(auth, authHandler)
	urlGroup := protected.Group("/urls")
	urlRoutes.RegisterRoutes(
	urlGroup,
	urlHandler,
)
	protected.GET("/profile", func(c *gin.Context) {

		userID := c.GetInt("userID")
		email := c.GetString("email")

		c.JSON(200, gin.H{
			"user_id": userID,
			"email":   email,
		})
	})
	return r
}
