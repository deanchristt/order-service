package entity

import "time"

type Order struct {
	ID           int            `gorm:"primary_key:auto_increment" json:"id"`
	Total        int            `gorm:"type:numeric" json:"total"`
	Number       string         `gorm:"type:varchar(50)" json:"number"`
	Date         time.Time      `json:"date"`
	ShippingDate time.Time      `json:"shippingDate"`
	IsDelivered  bool           `gorm:"not null" json:"isDelivered"`
	CustomerId   int            `gorm:"column:customer_id"`
	Customer     Customer       `json:"customer"`
	OrderProduct []OrderProduct `gorm:"many2many:order_product;"`
}
