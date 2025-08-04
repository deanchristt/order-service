package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JwtService interface {
	GenerateTokenUser(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtUserClaim struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func (j *jwtService) GenerateTokenUser(userId string) string {
	claims := &jwtUserClaim{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func NewJwtService() JwtService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    getIssuer(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	return secretKey
}

func getIssuer() string {
	issuer := os.Getenv("JWT_ISSUER")
	return issuer
}
