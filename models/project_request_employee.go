package models

import (
	"time"
)

type ProjectRequestEmployee struct {
	ID         int `gorm:"primary_key" json:"id"`
	ProjectId  int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project    Project
	EmployeeId int `gorm:"type:int; NOT NULL" json:"employee_id" binding:"required"`
	Employee   Employee
	Status     int       `gorm:"type:smallint; default:0; comment:0 = pending, 1 = approved" json:"status"`
	Noted      string    `gorm:"type:varchar(255);" json:"noted" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProjectRequestEmployees []ProjectRequestEmployee
