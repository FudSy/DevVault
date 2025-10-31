package config

import "github.com/FudSy/DevVault/internal/pkg/postgres"

type Config struct {
	Postgres postgres.Config
}

func InitCfg() (cfg *Config){
	return
}
