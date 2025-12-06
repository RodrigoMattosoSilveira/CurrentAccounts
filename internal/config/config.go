package config

import (
	"log/slog"
	"os"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV     string
	DB_NAME     string
	PROXY_PORT  string
	GIN_PORT    string
	FIBER_PORT  string
	CSRF_SECRET string
	SESSION_KEY string
	JWT_KEY	    string
}

var Cfg *Config

func LoadConfig() error {

	homeDir, err := utilities.FindProjectRoot()
	if (err != nil) {
        slog.Error("Error calculating project's home directory")		
	}

	envFile := homeDir + "/" + ".env"
    err = godotenv.Load(envFile)
    if err != nil {
        slog.Error("Error loading .env file")
		panic(err)
    }

	envSecretsFile := homeDir + "/" + ".env.secrets"
    err = godotenv.Load(envSecretsFile)
		if err != nil {
			slog.Error("Error loading .env.secrets file")
			panic(err)
		}

	Cfg = &Config{
		APP_ENV:     getEnv("APP_ENV", "development"),
		DB_NAME:     getEnv("DB_NAME", "/private/var/ContasCorrentes/sqlite_dev.db"),
		PROXY_PORT:  getEnv("PROXY_PORT", "80"),
		GIN_PORT:    getEnv("GIN_PORT", "8080"),
		FIBER_PORT:  getEnv("FIBER_PORT", "3000"),
		CSRF_SECRET: getEnv("CSRF_SECRET", "default-secret-must-be-32-chars-long"),
		SESSION_KEY: getEnv("SESSION_KEY", "default-secret-must-be-32-chars-long"),
		JWT_KEY:     getEnv("JWT_KEY",     "default-secret-must-be-32-chars-long"),
	}
	slog.Info("Configuration loaded", "app_env", Cfg.APP_ENV)
	return nil
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
