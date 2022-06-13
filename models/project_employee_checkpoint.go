package models

import (
	"time"
)

type ProjectEmployeeCheckpoint struct {
	ID                   int `gorm:"primary_key" json:"id"`
	ProjectId            int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project              *Project
	EmployeeId           int `gorm:"type:int; NOT NULL" json:"employee_id" binding:"required"`
	Employee             *Employee
	ProjectSectionRoomId int `gorm:"type:int; NOT NULL" json:"project_section_room_id" binding:"required"`
	ProjectSectionRoom   *ProjectSectionRoom
	QrId                 string    `gorm:"type:varchar(255);" json:"qr_id" binding:"required"`
	Photo                string    `gorm:"type:varchar(255);" json:"photo" binding:"required"`
	Noted                string    `gorm:"type:varchar(255);" json:"noted"`
	CreatedAt            time.Time `gorm:"NOT NULL" json:"created_at" binding:"required"`
}

type ProjectEmployeeCheckpoints []ProjectEmployeeCheckpoint
