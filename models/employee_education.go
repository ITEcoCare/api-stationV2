package models

import "time"

type EmployeeEducation struct {
	ID          int `gorm:"primary_key" json:"id"`
	EmployeeId  int `gorm:"type:int;NOT NULL" json:"employee_id"`
	Employee    *Employee
	Name        string    `gorm:"type:varchar(255);NOT NULL" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EmployeeEducations []EmployeeEducation
