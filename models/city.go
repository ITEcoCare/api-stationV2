package models

import "time"

type City struct {
	ID         int `gorm:"primary_key" json:"id"`
	ProvinceId int `gorm:"type:int;NOT NULL" json:"province_id"`
	Province   Province
	Name       string    `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Cities []City
