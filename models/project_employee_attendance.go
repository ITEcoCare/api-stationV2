package models

import (
	"time"
)

type ProjectEmployeeAttendace struct {
	ID                   int `gorm:"primary_key" json:"id"`
	ProjectId            int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project              Project
	EmployeeId           int `gorm:"type:int; NOT NULL" json:"employee_id" binding:"required"`
	Employee             Employee
	ProjectCustomShiftId int `gorm:"type:int; NOT NULL" json:"project_custom_shift_id" binding:"required"`
	ProjectCustomShift   ProjectCustomShift
	InOut                string    `gorm:"type:char(3); default:IN; comment:IN, OUT" json:"in_out" binding:"required"`
	Latitude             string    `gorm:"type:varchar(255); NOT NULL" json:"latitude" binding:"required"`
	Longitude            string    `gorm:"type:varchar(255); NOT NULL" json:"longitude" binding:"required"`
	CreatedAt            time.Time `gorm:"NOT NULL" json:"created_at" binding:"required"`
}

type ProjectEmployeeAttendaces []ProjectEmployeeAttendace
