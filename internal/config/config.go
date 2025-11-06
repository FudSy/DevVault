package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/FudSy/DevVault/internal/pkg/postgres"
)

type Config struct {
	Postgres postgres.Config
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
	}
	return cfg
}
