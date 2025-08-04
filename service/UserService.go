package service

import (
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/entity"
	"github.com/deanchristt/order-service/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	GetProfile(userId int) entity.User
	UpdateUser(dto dto.UserUpdate) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func (user userService) GetProfile(userId int) entity.User {
	return user.userRepository.ProfileUser(userId)
}

func (user userService) UpdateUser(dto dto.UserUpdate) entity.User {
	userUpdate := entity.User{}
	err := smapping.FillStruct(&userUpdate, smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updateCustomer := user.userRepository.UpdateUser(userUpdate)
	return updateCustomer
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRep,
	}
}
