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
	CreateUser(create dto.UserCreate) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func (a authService) VerifyCredential(email string, password string) interface{} {
	res := a.userRepository.VerifyCredential(email)
	if v, ok := res.(entity.User); ok {
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

func (a authService) CreateUser(create dto.UserCreate) entity.User {
	userCreate := entity.User{}
	err := smapping.FillStruct(&userCreate, smapping.MapFields(&create))
	if err != nil {
		log.Fatalf("Field map %v", err)
	}
	res := a.userRepository.InsertUser(userCreate)
	return res
}

func (a authService) FindByEmail(email string) entity.User {
	return a.userRepository.FindByEmail(email)
}

func (a authService) IsDuplicateEmail(email string) bool {
	res := a.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRep,
	}
}
