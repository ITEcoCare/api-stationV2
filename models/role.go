package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int            `gorm:"primary_key" json:"id"`
	CompanyId   int            `gorm:"type:int;NOT NULL" json:"company_id" binding:"required"`
	Company     *Company       `json:"-"`
	Name        string         `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Roles []Role
