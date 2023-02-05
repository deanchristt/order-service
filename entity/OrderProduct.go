package entity

type OrderProduct struct {
	ProductId int `gorm:"primaryKey" column:"product_id"`
	OrderId   int `gorm:"primaryKey" column:"order_id"`
	//Product   Product `gorm:"foreignkey:ID;references:ProductId""`
	//Order     Order   `gorm:"foreignkey:ID;references:OrderId"`
	Quantity int `gorm:"type:numeric" json:"quantity"`
}
