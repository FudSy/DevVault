package service

import "github.com/FudSy/DevVault/internal/pkg/postgres"

type Handlers struct {
	postgres *postgres.DB
}

func NewHandlers(postgres *postgres.DB) *Handlers {
	return &Handlers{
		postgres: postgres,
	}
}