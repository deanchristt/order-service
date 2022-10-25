package controller

import (
	"fmt"
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/helper"
	"github.com/deanchristt/order-service/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController interface {
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type customerController struct {
	customerService service.CustomerService
	jwtService      service.JwtService
}

func (c customerController) GetProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprintf("%v", claims["user_id"])
	customerId, _ := strconv.Atoi(userId)
	user := c.customerService.GetProfile(customerId)
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}

func (c customerController) UpdateProfile(ctx *gin.Context) {
	var customerUpdateDto dto.CustomerUpdate
	errDto := ctx.ShouldBind(&customerUpdateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		panic(err.Error())
	}
	customerUpdateDto.ID = id
	u := c.customerService.UpdateCustomer(customerUpdateDto)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}

func NewCustomerController(service service.CustomerService, jwtService service.JwtService) CustomerController {
	return &customerController{
		customerService: service,
		jwtService:      jwtService,
	}
}
