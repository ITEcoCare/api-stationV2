package models

import (
	"time"

	"gorm.io/gorm"
)

type JobPosition struct {
	ID          int `gorm:"primary_key" json:"id"`
	TeamId      int `gorm:"type:int;NOT NULL" json:"team_id" binding:"required"`
	Team        Team
	Code        string         `gorm:"type:varchar(50);NOT NULL;UNIQUE" json:"code" binding:"required"`
	Name        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	MinSalary   float64        `gorm:"NOT NULL" json:"min_salary"`
	MaxSalary   float64        `gorm:"NOT NULL" json:"max_salary"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type JobPositions []JobPosition
