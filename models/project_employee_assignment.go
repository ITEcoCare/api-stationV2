package models

import (
	"time"
)

type ProjectEmployeeAssignment struct {
	ID         int `gorm:"primary_key" json:"id"`
	ProjectId  int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project    *Project
	EmployeeId int `gorm:"type:int; NOT NULL" json:"employee_id" binding:"required"`
	Employee   *Employee
	Noted      string    `gorm:"type:varchar(255);" json:"noted" binding:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ProjectEmployeesAssignment []ProjectEmployeeAssignment
