package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	//Input service in here

}

// creates a new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (auth *authController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Login",
	})
}

func (auth *authController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Register",
	})
}
