package models

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID          int `gorm:"primary_key" json:"id"`
	BranchId    int `json:"branch_id" binding:"required"`
	Branch      *Branch
	Code        string         `gorm:"type:varchar(50);NOT NULL;UNIQUE" json:"code" binding:"required"`
	Name        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Warehouses []Warehouse
