package repository

import (
	"github.com/deanchristt/order-service/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository interface {
	InsertUser(customer entity.User) entity.User
	UpdateUser(customer entity.User) entity.User
	VerifyCredential(email string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(UserId int) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempCustomer entity.User
		db.connection.Find(&tempCustomer, user.ID)
		user.Password = tempCustomer.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string) interface{} {
	var customer entity.User
	res := db.connection.Where("email = ?", email).Take(&customer)
	if res.Error == nil {
		return customer
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var customer entity.User
	return db.connection.Where("email = ?", email).Take(&customer)
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var customer entity.User
	db.connection.Where("email = ?", email).Take(&customer)
	return customer
}

func (db *userConnection) ProfileUser(userID int) entity.User {
	var customer entity.User
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
