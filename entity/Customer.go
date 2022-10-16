package entity

import "gorm.io/gorm"

type Customer struct {
	ID       int    `gorm:"primary_key:auto_increment" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string `gorm:"->;<-;not null" json:"password"`
	Token    string `gorm:"-" json:"token,omitempty"`
	gorm.Model
}
