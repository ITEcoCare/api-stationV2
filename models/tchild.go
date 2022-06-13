package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Tchild struct {
	// ID          int            `gorm:"primary_key" json:"id"`
	// Tparent     *Tparent       `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID`
	// ParentID    uuid.UUID `gorm:"type:uuid REFERENCES tparents(id)"`
	ID          uuid.UUID      `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	ParentID    uuid.UUID      `gorm:"type:uuid REFERENCES tparents(id)" json:"parent_id"`
	Tparent     Tparent        `gorm:"ForeignKey:ParentID;AssociationForeignKey:ID"`
	Name        string         `gorm:"type:varchar(255);NOT NULL;UNIQUE" json:"name" binding:"required"`
	Description string         `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"`
}

type Tchilds []Tchild
