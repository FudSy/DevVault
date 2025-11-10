package postgres

import (
	"errors"
	"fmt"

	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/rs/zerolog/log"
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

func New(dsn string) (*DB) {
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("Error while opening connection to DB")
	}
	log.Info().Msg("DB connection opened")
	db := &DB{Database: gdb}
	return db
}

func (db *DB) Migrate() {
	err := db.Database.AutoMigrate(&models.User{}, &models.Snippet{}, &models.Favorite{})
	if err != nil {
		log.Fatal().Msg("Error while migrating DB")
	}
}

// User 
func (d *DB) CreateUser(user *models.User) error{
	return d.Database.Create(&user).Error
}

func (d *DB) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	err := d.Database.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Snippet
func (d *DB) CreateSnippet(snippet models.Snippet) error {
	return d.Database.Create(&snippet).Error
}