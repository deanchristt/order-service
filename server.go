package main

import (
	"github.com/deanchristt/order-service/config"
	"github.com/deanchristt/order-service/controller"
	"github.com/deanchristt/order-service/repository"
	"github.com/deanchristt/order-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                 *gorm.DB                      = config.SetupDatabaseConnection()
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
	jwtService         service.JwtService            = service.NewJwtService()
	authService        service.AuthService           = service.NewAuthService(customerRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
