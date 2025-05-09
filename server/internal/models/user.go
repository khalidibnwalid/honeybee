package models

import (
	"database/sql"
	"khalidibnwalid/luma_server/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Username       string         `gorm:"unique;type:varchar(255)" json:"username"`
	Email          string         `gorm:"unique;type:varchar(255)" json:"email"`
	HashedPassword string         `gorm:"type:text" json:"-"`
	AvatarURL      sql.NullString `gorm:"type:text" json:"avatarUrl"`
	CreatedAt      time.Time      `gorm:"default:now()" json:"createdAt"`
	UpdatedAt      time.Time      `gorm:"default:now()" json:"updatedAt"`
}

func (User) TableName() string {
	return "luma.users"
}

// # User model methods

// # Database operations

func (u *User) Create(db *database.Database) error {
	return db.Client.Create(u).Error
}

func (u *User) Update(db *database.Database) error {
	return db.Client.Updates(u).Error
}

func (u *User) Delete(db *database.Database) error {
	return db.Client.Delete(u).Error
}

func (u *User) GetByID(db *database.Database, id uuid.UUID) error {
	return db.Client.First(u, "id = ?", id).Error
}

func (u *User) GetByUsername(db *database.Database, username string) error {
	return db.Client.First(u, "username = ?", username).Error
}

// func (u *User) GetByEmail(db *database.Database, email mail.Address) error {
// 	return db.Client.First(u, "email = ?", email.Address).Error
// }
