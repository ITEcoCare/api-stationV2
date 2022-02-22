package models

import (
	"time"
)

type ProjectCustomShift struct {
	ID        int `gorm:"primary_key" json:"id"`
	ProjectId int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project   *Project
	Name      string    `gorm:"type:varchar(255);NOT NULL" json:"name" binding:"required"`
	HourIn    time.Time `gorm:"type:time;NOT NULL" json:"hour_in" binding:"required"`
	HourOut   time.Time `gorm:"type:time;NOT NULL" json:"hour_out" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectCustomShifts []ProjectCustomShift
