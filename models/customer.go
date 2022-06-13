package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID              int `gorm:"primary_key" json:"id"`
	BranchId        int `json:"branch_id" binding:"required"`
	Branch          *Branch
	UserId          int `json:"user_id" binding:"required"`
	User            *User
	CustomerGroupId int `gorm:"type:int;NOT NULL" json:"customer_group_id" binding:"required"`
	CustomerGroup   *CustomerGroup
	ProvinceId      int `gorm:"type:int;NOT NULL" json:"province_id" binding:"required"`
	Province        *Province
	BillToAddress   int            `gorm:"type:int;NOT NULL" json:"bill_to_address" binding:"required"`
	Name            string         `gorm:"type:varchar(255);" json:"name" binding:"required"`
	Status          int            `gorm:"type:smallint; default:1; comment:0 = inactive, 1 = active" json:"status"`
	Phone           string         `gorm:"type:varchar(50)" json:"phone" binding:"required"`
	Email           string         `gorm:"type:varchar(50)" json:"email" binding:"required"`
	Noted           string         `gorm:"type:varchar(255)" json:"noted" binding:"required"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-"`
}

type Customers []Customer
