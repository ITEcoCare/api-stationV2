package models

import (
	"time"
)

type ProjectTop struct {
	ID        int `gorm:"primary_key" json:"id"`
	ProjectId int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project   Project
	Percent   int       `gorm:"type:smallint;" json:"percent"`
	Noted     string    `gorm:"type:varchar(255);" json:"noted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectsTop []ProjectTop
