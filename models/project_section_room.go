package models

import (
	"time"
)

type ProjectSectionRoom struct {
	ID        int `gorm:"primary_key" json:"id"`
	ProjectId int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project   *Project
	Area      string    `gorm:"type:varchar(255);NOT NULL" json:"area" binding:"required"`
	Name      string    `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	Noted     string    `gorm:"type:varchar(255);" json:"noted" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectSectionRooms []ProjectSectionRoom
