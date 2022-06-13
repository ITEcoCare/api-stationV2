package models

import (
	"time"

	"gorm.io/gorm"
)

type Branch struct {
	ID          int `gorm:"primary_key" json:"id"`
	CompanyId   int `gorm:"type:int;NOT NULL" json:"company_id" binding:"required"`
	Company     *Company
	ProvinceId  int `gorm:"type:int;NOT NULL" json:"province_id" binding:"required"`
	Province    *Province
	CityId      int `gorm:"type:int;NOT NULL" json:"city_id" binding:"required"`
	City        *City
	Code        string         `gorm:"type:varchar(50);NOT NULL;UNIQUE" json:"code" binding:"required"`
	Name        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Branches []Branch
