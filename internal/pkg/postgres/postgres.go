package postgres

import (
	"errors"
	"fmt"

	"github.com/FudSy/DevVault/internal/pkg/models"
	"github.com/google/uuid"
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

func (d *DB) UpdateSnippet(snippet models.Snippet, id uuid.UUID, userID uuid.UUID) error {
	var result models.Snippet

	err := d.Database.Where("id = ? AND user_id = ?", id, userID).First(&result).Error
	if err != nil {
		return err
	}

	return d.Database.Model(&result).Updates(map[string]interface{}{
		"title":       snippet.Title,
		"code":        snippet.Code,
		"language":    snippet.Language,
		"description": snippet.Description,
		"is_public":   snippet.IsPublic,
		"updated_at": snippet.UpdatedAt,
	}).Error
}

func (d *DB) DeleteSnippet(id uuid.UUID, userID uuid.UUID) error {
	var result models.Snippet

	err := d.Database.Where("id = ? AND user_id = ?", id, userID).First(&result).Error
	if err != nil {
		return err
	}

	return d.Database.Delete(&result).Error
}

func (d *DB) GetSnippet(id uuid.UUID, userID uuid.UUID) (models.Snippet, error) {
	var result models.Snippet

	err := d.Database.Where("id = ?", id).First(&result).Error
	if err != nil {
		return models.Snippet{}, err
	}
	if !result.IsPublic && result.UserID != userID {
		return models.Snippet{}, errors.New("not enough permissions to view")
	}

	return result, nil
}