package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID            int `gorm:"primary_key" json:"id"`
	BranchId      int `gorm:"type:int;NOT NULL" json:"branch_id" binding:"required"`
	Branch        Branch
	CustomerId    int `gorm:"type:int;NOT NULL" json:"customer_id" binding:"required"`
	Customer      Customer
	Status        int            `gorm:"type:smallint; default:0; comment:0 = draft, 1 = complete, 2 = on_going, 3 = terminate, 4 = out_of_date" json:"status"`
	ProjectName   string         `gorm:"type:varchar(200); NOT NULL" json:"project_name" binding:"required"`
	ProjectNumber string         `gorm:"type:varchar(200); NOT NULL" json:"project_number" binding:"required"`
	ProjectValue  float64        `gorm:"NOT NULL" json:"project_value" binding:"required"`
	PriodStart    *time.Time     `gorm:"NOT NULL" json:"period_start"`
	PeriodEnd     *time.Time     `gorm:"NOT NULL" json:"period_end"`
	Noted         string         `gorm:"type:varchar(255);NOT NULL" json:"noted"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

type Projects []Project
