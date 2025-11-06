package config

import (
	"log"
	"os"
	"strings"

	"github.com/FudSy/DevVault/internal/pkg/postgres"
	logger "github.com/FudSy/DevVault/pkg"
	"github.com/joho/godotenv"
)

type Config struct {
	Postgres postgres.Config
	Logger logger.Config
}

func InitCfg() (cfg *Config){
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	cfg = &Config{
		Postgres: postgres.Config{
			DBName: os.Getenv("PG_DBName"),
			Host: os.Getenv("PG_Host"),
			Password: os.Getenv("PG_Password"),
			Port: os.Getenv("PG_Port"),
			SSLMode: os.Getenv("PG_SSLMode"),
			User: os.Getenv("PG_User"),
		},
		Logger: logger.Config{
			Level: strings.ToLower(os.Getenv("LOG_LEVEL")),
		},
	}
	return cfg
}
