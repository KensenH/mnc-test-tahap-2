package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort   string
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
	JwtSecret  string
}

func New() (Config, error) {
	var err error
	var cfg Config

	err = godotenv.Load(".env")
	if err != nil {
		return cfg, err
	}

	cfg.HttpPort = getOrDefaultEnv("HTTP_PORT", ":8080")
	cfg.DbHost = getOrDefaultEnv("DB_HOST", ":5432")
	cfg.DbPort = getOrDefaultEnv("DB_PORT", "localhost")
	cfg.DbName = getOrDefaultEnv("DB_NAME", "postgres")
	cfg.DbUser = getOrDefaultEnv("DB_USER", "postgres")
	cfg.DbPassword = getOrDefaultEnv("DB_PASSWORD", "")
	cfg.JwtSecret = getOrDefaultEnv("JWT_SECRET", "")

	return cfg, err
}

func getOrDefaultEnv(envName string, def string) string {
	value := os.Getenv(envName)
	if value == "" {
		return def
	}

	return value
}
