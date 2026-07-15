package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret string

	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {

	_ = godotenv.Load(".env.local")
	return &Config{
		Port: os.Getenv("PORT"),

		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),

		JWTSecret: os.Getenv("JWT_SECRET"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
	}
}
