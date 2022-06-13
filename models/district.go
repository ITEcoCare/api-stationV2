package models

import "time"

type District struct {
	ID        int `gorm:"primary_key" json:"id"`
	CityId    int `gorm:"type:int;NOT NULL" json:"city_id"`
	City      City
	Name      string    `gorm:"type:varchar(255);NOT NULL;" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Districts []District
