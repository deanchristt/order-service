package entity

import "time"

type Transaction struct {
	ID                int       `gorm:"primary_key:auto_increment" json:"id"`
	ReferenceNumber   string    `gorm:"type:varchar(50)" json:"referenceNumber"`
	TransactionNumber string    `gorm:"type:varchar(50)" json:"transactionNumber"`
	CreatedAt         time.Time `json:"createdAt"`
	PaidAt            time.Time `json:"paidAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	CustomerId        int       `gorm:"not null"`
	Customer          Customer  `gorm:"foreignkey:CustomerId" json:"customer"`
	Payment           []Payment `json:"payment"`
	OrderId           int       `gorm:"ForeignKey:id;column:order_id"` //this foreignKey tag didn't works
	Order             Order     `gorm:"foreignKey:OrderId" json:"order"`
}
