package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Language string

const (
	Golang Language = "Golang"
	Python Language = "Python"
	CS Language = "C#"
	CPP Language = "C++"
	Java Language = "Java"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string     `gorm:"not null;unique" json:"username"`
	Email     string     `gorm:"not null;unique" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`

	Snippets  []Snippet
	Favorites []Favorite
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

type Snippet struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string     `gorm:"not null" json:"title"`
	Code        string     `gorm:"type:text;not null" json:"code"`
	Language    Language   `gorm:"type:varchar(20);not null" json:"language"`
	Description string     `json:"description"`
	IsPublic    bool       `gorm:"default:false" json:"is_public"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	User User `gorm:"foreignKey:UserID"`
}

func (s *Snippet) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}

type Favorite struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	SnippetID uuid.UUID  `gorm:"type:uuid;not null;index" json:"snippet_id"`

	User    User    `gorm:"foreignKey:UserID"`
	Snippet Snippet `gorm:"foreignKey:SnippetID"`
}


func (f *Favorite) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}