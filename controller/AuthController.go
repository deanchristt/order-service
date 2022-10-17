package controller

import (
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/entity"
	"github.com/deanchristt/order-service/helper"
	"github.com/deanchristt/order-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	//Input service in here
	authService service.AuthService
	jwtService  service.JwtService
}

// creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (auth *authController) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	errDto := c.ShouldBind(&loginDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := auth.authService.VerifiyCredential(loginDto.Email, loginDto.Password)
	if v, ok := authResult.(entity.Customer); ok {
		convertId := strconv.FormatInt(int64(v.ID), 10)
		generatedToken := auth.jwtService.GenerateToken(convertId)
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		c.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (auth *authController) Register(c *gin.Context) {
	var registerDto dto.CustomerCreate
	errDto := c.ShouldBind(&registerDto)
	if errDto != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !auth.authService.IsDuplicateEmail(registerDto.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		c.JSON(http.StatusConflict, response)
	} else {
		createdCustomer := auth.authService.CreateCustomer(registerDto)
		token := auth.jwtService.GenerateToken(strconv.FormatInt(int64(createdCustomer.ID), 10))
		createdCustomer.Token = token
		response := helper.BuildResponse(true, "OK!", createdCustomer)
		c.JSON(http.StatusCreated, response)
	}
}
