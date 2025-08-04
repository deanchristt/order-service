package entity

type OrderProduct struct {
	ProductId int `gorm:"primaryKey" column:"product_id"`
	OrderId   int `gorm:"primaryKey" column:"order_id"`
	Quantity  int `gorm:"type:numeric" json:"quantity"`
}
