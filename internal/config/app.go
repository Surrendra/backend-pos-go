package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	DBTimeZone string

	DBTimezone string

	AppEnv   string
	AppDebug bool
	AppPort  string

	JwtSecret string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisUsername string
}

func Load() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBTimeZone: os.Getenv("DB_TIMEZONE"),

		AppPort:  os.Getenv("APP_PORT"),
		AppEnv:   os.Getenv("APP_ENV"),
		AppDebug: os.Getenv("APP_DEBUG") == "true",

		JwtSecret: os.Getenv("JWT_SECRET"),

		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisUsername: os.Getenv("REDIS_USERNAME"),
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}
	if cfg.DBHost == "" || cfg.DBName == "" {
		log.Fatal("Database configuration is missing")
	}
	if cfg.DBTimeZone == "" {
		cfg.DBTimeZone = "Asia/Makkasar"
	}
	return cfg
}
