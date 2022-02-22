package models

import (
	"time"
)

type EmployeeAddress struct {
	ID         int `gorm:"primary_key" json:"id"`
	EmployeeId int `gorm:"type:int;NOT NULL" json:"employee_id"`
	Employee   *Employee
	ProvinceId int `gorm:"type:int;NOT NULL" json:"province_id"`
	Province   *Province
	CityId     int `gorm:"type:int;NOT NULL" json:"city_id"`
	City       *City
	DistrictId int `gorm:"type:int;NOT NULL" json:"district_id"`
	District   *District
	Address    string    `gorm:"type:text;" json:"address"`
	Latitude   string    `gorm:"type:varchar(200);" json:"latitude"`
	Longitude  string    `gorm:"type:varchar(200);" json:"longitude"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmployeeAddresses []EmployeeAddress
