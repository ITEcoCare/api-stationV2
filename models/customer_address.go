package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerAddress struct {
	ID         int `gorm:"primary_key" json:"id"`
	CustomerId int `gorm:"type:int;NOT NULL" json:"customer_id"`
	Customer   *Customer
	ProvinceId int `gorm:"type:int;NOT NULL" json:"province_id"`
	Province   *Province
	CityId     int `gorm:"type:int;NOT NULL" json:"city_id"`
	City       *City
	DistrictId int `gorm:"type:int;NOT NULL" json:"district_id"`
	District   *District
	Address    string         `gorm:"type:text;" json:"address"`
	Latitude   string         `gorm:"type:varchar(200);" json:"latitude"`
	Longitude  string         `gorm:"type:varchar(200);" json:"longitude"`
	IsPrimary  bool           `gorm:"type:boolean;" json:"is_primary"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

type CustomerAddresses []CustomerAddress
