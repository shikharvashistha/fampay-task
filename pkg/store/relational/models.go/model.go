package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrRequiredFieldNotPresent = errors.New("required field not present")
)

func RegisterSchema(db *gorm.DB) error {
	return db.AutoMigrate(
		&Videos{},
		&Cron{},
	)
}

type Model struct {
	ID        string    `gorm:"size:191; primaryKey;" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
