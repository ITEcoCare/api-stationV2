package models

import (
	"time"
)

type ProjectInvoice struct {
	ID             int `gorm:"primary_key" json:"id"`
	ProjectId      int `gorm:"type:int; NOT NULL" json:"project_id" binding:"required"`
	Project        Project
	ProjectTopId   int `gorm:"type:int; NOT NULL" json:"project_top_id" binding:"required"`
	CustomerId     int `gorm:"type:int; NOT NULL" json:"customer_id" binding:"required"`
	Customer       Customer
	InvoiceNumber  string     `gorm:"type:varchar(255); NOT NULL" json:"invoice_number" binding:"required"`
	InvoiceDueDate *time.Time `gorm:"NOT NULL" json:"invoice_due_date" binding:"required"`
	Amount         float64    `gorm:"NOT NULL;" json:"amount"`
	Noted          string     `gorm:"type:varchar(255);" json:"noted"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type ProjectInvoices []ProjectInvoice
