package postgres

import (
	"github.com/FudSy/DevVault/internal/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	User string
	Password string
	DBName string
	Port string
	SSLMode string
}


func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Snippet{}, &models.Favorite{})
}

