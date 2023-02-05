package entity

import "github.com/shopspring/decimal"

type Product struct {
	ID          int             `gorm:"primary_key:auto_increment" json:"id"`
	Name        string          `gorm:"type:varchar(255)" json:"name"`
	Description string          `gorm:"type:text" json:"description"`
	Image1      string          `gorm:"type:varchar(255)" json:"image1"`
	Image2      string          `gorm:"type:varchar(255)" json:"image2"`
	Image3      string          `gorm:"type:varchar(255)" json:"image3"`
	Price       decimal.Decimal `gorm:"type:decimal(19,2)" json:"price"`
	Stock       int             `gorm:"type:numeric" json:"stock"`
	SellerId    int             `gorm:"column:seller_id"`
	Seller      Seller          `json:"seller"`
}
