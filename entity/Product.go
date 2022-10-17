package entity

type Product struct {
	ID          int      `gorm:"primary_key:auto_increment" json:"id"`
	Title       string   `gorm:"type:varchar(255)" json:"title"`
	Description string   `gorm:"type:text" json:"description"`
	CustomerId  int      `gorm:"not null" json:"-"`
	Customer    Customer `gorm:"foreignkey:CustomerId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"customer"`
}
