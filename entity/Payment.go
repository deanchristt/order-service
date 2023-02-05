package entity

import "time"

type Payment struct {
	ID            int         `gorm:"primary_key:auto_increment" json:"id"`
	PaymentMethod string      `gorm:"type:varchar(50)" json:"referenceNumber"`
	Amount        string      `gorm:"type:varchar(50)" json:"transactionNumber"`
	PaidAt        time.Time   `json:"paidAt"`
	TransactionId int         `gorm:"not null"`
	Transaction   Transaction `gorm:"foreignkey:TransactionId" json:"transaction"`
}
