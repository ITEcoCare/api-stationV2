package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int `gorm:"primary_key" json:"id"`
	RoleId          int `gorm:"type:int;NOT NULL" json:"role_id" binding:"required"`
	Role            *Role
	Name            string         `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Username        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"username" binding:"required"`
	Email           string         `gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE_INDEX" json:"email" binding:"required"`
	Password        string         `gorm:"type:varchar(255)" json:"password"`
	Avatar          string         `gorm:"type:varchar(255)" json:"avatar"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

type Users []User
