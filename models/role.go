package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int            `gorm:"primary_key" json:"id"`
	Name        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Roles []Role
