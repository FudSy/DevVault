package main

import (
	"github.com/FudSy/DevVault/internal/config"
	"github.com/FudSy/DevVault/internal/pkg/postgres"
	"github.com/FudSy/DevVault/internal/pkg/service"
)

func main() {
	cfg := config.InitCfg()

	cfg.Logger.Init()

	db := postgres.New(cfg.Postgres.DSN())
	db.Migrate()

	r := service.Router(db)

	r.Run()
}