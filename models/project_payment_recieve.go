package models

import (
	"time"
)

type ProjectPaymentReceive struct {
	ID               int `gorm:"primary_key" json:"id"`
	ProjectInvoiceId int `gorm:"type:int; NOT NULL" json:"project_invoice_id" binding:"required"`
	ProjectInvoice   ProjectInvoice
	PayAmount        float64   `gorm:"NOT NULL;" json:"pay_amount"`
	Noted            string    `gorm:"type:varchar(255);" json:"noted"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ProjectPaymentReceives []ProjectPaymentReceive
