package controller

import (
	"fmt"
	"github.com/deanchristt/order-service/dto"
	"github.com/deanchristt/order-service/entity"
	"github.com/deanchristt/order-service/helper"
	"github.com/deanchristt/order-service/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService service.ProductService
	jwtService     service.JwtService
}

func (p *productController) getCustomerIdByToken(token string) string {
	aToken, err := p.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (p productController) All(context *gin.Context) {
	var products = p.productService.All()
	res := helper.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)
}

func (p productController) FindById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		res := helper.BuildErrorResponse("Param id was not found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var product = p.productService.FindById(id)
	if (product == entity.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", product)
		context.JSON(http.StatusOK, res)
	}
}

func (p productController) Insert(context *gin.Context) {
	var productCreatDto dto.ProductCreateDto
	errDto := context.ShouldBind(&productCreatDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		customerId := p.getCustomerIdByToken(authHeader)
		convertCustomerId, err := strconv.Atoi(customerId)
		if err == nil {
			productCreatDto.CustomerId = convertCustomerId
		}
		result := p.productService.Insert(productCreatDto)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (p productController) Update(context *gin.Context) {
	var productUpdateDto dto.ProductUpdateDto
	errDto := context.ShouldBind(&productUpdateDto)
	if errDto != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDto.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := p.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	customerId := fmt.Sprintf("%v", claims["user_id"])
	path := context.Param("id")
	productId, _ := strconv.Atoi(path)
	if p.productService.IsAllowedToEdit(customerId, productId) {
		id, errID := strconv.Atoi(customerId)
		if errID == nil {
			productUpdateDto.ID = productId
			productUpdateDto.CustomerId = id
		}
		result := p.productService.Update(productUpdateDto)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (p productController) Delete(context *gin.Context) {
	//TODO implement me
	var product entity.Product
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	product.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := p.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if p.productService.IsAllowedToEdit(userID, product.ID) {
		p.productService.Delete(product)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func NewProductController(productService service.ProductService, jwtService service.JwtService) ProductController {
	return &productController{
		productService: productService,
		jwtService:     jwtService,
	}
}
