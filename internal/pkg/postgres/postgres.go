package postgres

import (
	"fmt"
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

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func New(dsn string) (*DB, error) {
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db := &DB{Database: gdb}
	return db, err
}

func (db *DB) Migrate() {
	db.Database.AutoMigrate(&models.User{}, &models.Snippet{}, &models.Favorite{})
}

func (d *DB) CreateUser(user *models.User) error{
	return d.Database.Create(&user).Error
}

func (d *DB) GetUser(id uuid.UUID) error{
	return d.Database.Find(&models.User{}).Where("id = ?", id).Error
}
