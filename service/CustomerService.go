package service

import (
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/entity"
	"github.com/deanchristt/order-service/repository"
	"github.com/mashingan/smapping"
	"log"
)

type CustomerService interface {
	GetProfile(userId int) entity.Customer
	UpdateCustomer(dto dto.CustomerUpdate) entity.Customer
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func (c customerService) GetProfile(userId int) entity.Customer {
	return c.customerRepository.ProfileUser(userId)
}

func (c customerService) UpdateCustomer(dto dto.CustomerUpdate) entity.Customer {
	userUpdate := entity.Customer{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateCustomer := c.customerRepository.UpdateCustomer(userUpdate)
	return updateCustomer
}

func NewCustomerService(customerRep repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRep,
	}
}
