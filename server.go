package main

import (
	"github.com/deanchristt/order-service/config"
	"github.com/deanchristt/order-service/controller"
	"github.com/deanchristt/order-service/repository"
	"github.com/deanchristt/order-service/routes"
	"github.com/deanchristt/order-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	jwtService        service.JwtService           = service.NewJwtService()
	authService       service.AuthService          = service.NewAuthService(userRepository)
	userService       service.UserService          = service.NewUserService(userRepository)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	productService    service.ProductService       = service.NewProductService(productRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	productController controller.ProductController = controller.NewProductController(productService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	routes.RegisterRoutes(r, authController, userController, jwtService, productController)

	r.Run()
}
