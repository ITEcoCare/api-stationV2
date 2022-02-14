package models

import "time"

type Province struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Provinces []Province
