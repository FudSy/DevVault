package postgres

import (
	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/google/uuid"
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

type DB struct {
	Database *gorm.DB
}
	

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Snippet{}, &models.Favorite{})
}

func (d *DB) CreateUser(user *models.User) error{
	return d.Database.Create(&user).Error
}

func (d *DB) GetUser(id uuid.UUID) error{
	return d.Database.Find(&models.User{}).Where("id = ?", id).Error
}
