package database

import (
	"context"
	"fmt"
	"log"

	"github.com/AbhayBajpai10/url-shortener/config"

	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) *pgx.Conn {

	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to PostgreSQL")

	return conn
}