package entity

type Customer struct {
	ID          int           `gorm:"primary_key:auto_increment" json:"id"`
	FirstName   string        `gorm:"type:varchar(50)" json:"firstName"`
	LastName    string        `gorm:"type:varchar(50)" json:"lastName"`
	Email       string        `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password    string        `gorm:"->;<-;not null" json:"-"`
	PhoneNumber string        `gorm:"type:varchar(20)" json:"phoneNumber"`
	Address     string        `gorm:"type:varchar(255)" json:"address"`
	Token       string        `gorm:"-" json:"token,omitempty"`
	Transaction []Transaction `gorm:"foreignKey:CustomerId" json:"transaction"`
	Order       []Order       `gorm:"foreignKey:CustomerId" json:"order"`
}
