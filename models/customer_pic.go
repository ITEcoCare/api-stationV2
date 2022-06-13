package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerPic struct {
	ID         int `gorm:"primary_key" json:"id"`
	CustomerId int `gorm:"type:int;NOT NULL" json:"customer_id" binding:"required"`
	Customer   *Customer
	ProvinceId int `gorm:"type:int;NOT NULL" json:"province_id" binding:"required"`
	Province   *Province
	CityId     int `gorm:"type:int;NOT NULL" json:"city_id" binding:"required"`
	City       *City
	DistrictId int            `gorm:"type:int;NOT NULL" json:"district_id" binding:"required"`
	Name       string         `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Address    string         `gorm:"type:text; NOT NULL" json:"address"`
	Phone_01   string         `gorm:"type:varchar(30); NOT NULL" json:"phone_01"`
	Phone_02   string         `gorm:"type:varchar(30); " json:"phone_02"`
	Email      string         `json:"email"`
	Noted      string         `json:"noted"`
	Status     int            `gorm:"type:smallint; default:1; comment:0 = inactive, 1 = active" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

type CustomerPics []CustomerPic
