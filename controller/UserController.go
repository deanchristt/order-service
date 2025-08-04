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

type UserController interface {
	GetProfile(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JwtService
}

func (c userController) GetProfile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userIDStr := fmt.Sprintf("%v", claims["userId"])
	userId, err := strconv.Atoi(userIDStr)
	user := c.userService.GetProfile(userId)
	res := helper.BuildResponse(true, "OK", user)
	ctx.JSON(http.StatusOK, res)
}

func (c userController) UpdateProfile(ctx *gin.Context) {
	var userUpdateDto dto.UserUpdate
	errDto := ctx.ShouldBind(&userUpdateDto)
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
	id, err := strconv.Atoi(fmt.Sprintf("%v", claims["userId"]))
	if err != nil {
		panic(err.Error())
	}
	userUpdateDto.ID = id
	u := c.userService.UpdateUser(userUpdateDto)
	res := helper.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}

func NewUserController(service service.UserService, jwtService service.JwtService) UserController {
	return &userController{
		userService: service,
		jwtService:  jwtService,
	}
}
