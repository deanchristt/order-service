package repository

import (
	"github.com/deanchristt/order-service/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type CustomerRepository interface {
	InsertCustomer(customer entity.Customer) entity.Customer
	UpdateCustomer(customer entity.Customer) entity.Customer
	VerifyCredential(email string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.Customer
	ProfileUser(UserId int) entity.Customer
}

type customerConnection struct {
	connection *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerConnection{
		connection: db,
	}
}

func (db *customerConnection) InsertCustomer(customer entity.Customer) entity.Customer {
	customer.Password = hashAndSalt([]byte(customer.Password))
	db.connection.Save(&customer)
	return customer
}

func (db *customerConnection) UpdateCustomer(customer entity.Customer) entity.Customer {
	if customer.Password != "" {
		customer.Password = hashAndSalt([]byte(customer.Password))
	} else {
		var tempCustomer entity.Customer
		db.connection.Find(&tempCustomer, customer.ID)
		customer.Password = tempCustomer.Password
	}
	db.connection.Save(&customer)
	return customer
}

func (db *customerConnection) VerifyCredential(email string) interface{} {
	var customer entity.Customer
	res := db.connection.Where("email = ?", email).Take(&customer)
	if res.Error == nil {
		return customer
	}
	return nil
}

func (db *customerConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var customer entity.Customer
	return db.connection.Where("email = ?", email).Take(&customer)
}

func (db *customerConnection) FindByEmail(email string) entity.Customer {
	var customer entity.Customer
	db.connection.Where("email = ?", email).Take(&customer)
	return customer
}

func (db *customerConnection) ProfileUser(userID int) entity.Customer {
	var customer entity.Customer
	db.connection.Find(&customer, userID)
	return customer
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
