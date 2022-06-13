package models

import "time"

type Village struct {
	ID         int `gorm:"primary_key" json:"id"`
	DistrictId int `gorm:"type:int;NOT NULL" json:"district_id"`
	District   District
	Name       string    `gorm:"type:varchar(255);NOT NULL;" json:"name" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Villages []Village
