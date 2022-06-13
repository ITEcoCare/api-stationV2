package models

import "time"

type EmployeeContact struct {
	ID         int `gorm:"primary_key" json:"id"`
	EmployeeId int `gorm:"type:int;NOT NULL" json:"employee_id"`
	Employee   *Employee
	Phone      string    `gorm:"type:varchar(50);NOT NULL" json:"phone"`
	Email      string    `gorm:"type:varchar(50);NOT NULL" json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EmployeeContacts []EmployeeContact
