package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID             int `gorm:"primary_key" json:"id"`
	BranchId       int `json:"branch_id" binding:"required"`
	Branch         *Branch
	UserId         int `json:"user_id" binding:"required"`
	User           *User
	JobPositionId  int `json:"job_position_id" binding:"required"`
	JobPosition    *JobPosition
	IdentityCardId int `json:"identity_card_id"`
	IdentityCard   *IdentityCard
	Nik            string         `gorm:"type:varchar(100);NOT NULL;UNIQUE" json:"nik" binding:"required"`
	Firstname      string         `gorm:"type:varchar(255);NOT NULL;" json:"firstname" binding:"required"`
	Lastname       string         `gorm:"type:varchar(255);" json:"lastname"`
	Avatar         string         `gorm:"type:varchar(255);" json:"avatar"`
	BirthPlace     string         `gorm:"type:varchar(255);" json:"birth_place"`
	BirthDate      *time.Time     `gorm:"type:date;" json:"birth_date"`
	MaritalStatus  string         `gorm:"type:char(20);" json:"marital_status"`
	Gender         string         `gorm:"type:char(10);" json:"gender"`
	HireDate       *time.Time     `gorm:"type:date;" json:"hire_date"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}

type Employees []Employee
