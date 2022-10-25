package service

import (
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/entity"
	"github.com/deanchristt/order-service/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateCustomer(create dto.CustomerCreate) entity.Customer
	FindByEmail(email string) entity.Customer
	IsDuplicateEmail(email string) bool
}

type authService struct {
	customerRepository repository.CustomerRepository
}

func (a authService) VerifyCredential(email string, password string) interface{} {
	res := a.customerRepository.VerifyCredential(email)
	if v, ok := res.(entity.Customer); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (a authService) CreateCustomer(create dto.CustomerCreate) entity.Customer {
	customerCreate := entity.Customer{}
	err := smapping.FillStruct(&customerCreate, smapping.MapFields(&create))
	if err != nil {
		log.Fatalf("Field map %v", err)
	}
	res := a.customerRepository.InsertCustomer(customerCreate)
	return res
}

func (a authService) FindByEmail(email string) entity.Customer {
	return a.customerRepository.FindByEmail(email)
}

func (a authService) IsDuplicateEmail(email string) bool {
	res := a.customerRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(customerRep repository.CustomerRepository) AuthService {
	return &authService{
		customerRep,
	}
}
