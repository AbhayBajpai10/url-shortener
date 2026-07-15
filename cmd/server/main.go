package main

import (
	"context"
	"fmt"

	"github.com/AbhayBajpai10/url-shortener/config"

	authHandler "github.com/AbhayBajpai10/url-shortener/internal/auth/handler"
	authRepository "github.com/AbhayBajpai10/url-shortener/internal/auth/repository"
	authService "github.com/AbhayBajpai10/url-shortener/internal/auth/service"

	urlHandler "github.com/AbhayBajpai10/url-shortener/internal/url/handler"
	urlRepository "github.com/AbhayBajpai10/url-shortener/internal/url/repository"
	urlService "github.com/AbhayBajpai10/url-shortener/internal/url/service"

	"github.com/AbhayBajpai10/url-shortener/internal/database"
	"github.com/AbhayBajpai10/url-shortener/internal/router"
	"github.com/AbhayBajpai10/url-shortener/internal/cache"
)

func main() {

	cfg := config.LoadConfig()

	conn := database.Connect(cfg)
	redisClient:=cache.ConnectRedis(cfg)
	defer conn.Close(context.Background())
	// Auth

	repo := authRepository.NewRepository(conn)
	authService := authService.NewService(repo, cfg)
	authHandler := authHandler.NewHandler(authService)

	// URL

	urlRepo := urlRepository.NewRepository(conn,redisClient)
	urlService := urlService.NewService(urlRepo)
	urlHandler := urlHandler.NewHandler(urlService)

	r := router.SetupRouter(
		authHandler,
		urlHandler,
		cfg,
	)

	fmt.Println("Server started on port", cfg.Port)

	r.Run(":" + cfg.Port)
}
