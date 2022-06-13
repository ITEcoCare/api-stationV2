package models

import "time"

type EmployeeAttendance struct {
	ID         int `gorm:"primary_key" json:"id"`
	EmployeeId int `gorm:"type:int;NOT NULL" json:"employee_id"`
	Employee   *Employee
	BranchId   int `gorm:"type:int;NOT NULL" json:"branch_id"`
	Branch     *Branch
	Timedate   *time.Time `gorm:"NOT NULL" json:"timedate"`
	InOut      string     `gorm:"type:char(10);NOT NULL" json:"in_out"`
	CreatedAt  time.Time  `json:"created_at"`
}

type EmployeeAttendances []EmployeeAttendance
