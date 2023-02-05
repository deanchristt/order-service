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
	productRepository  repository.ProductRepository  = repository.NewProductRepository(db)
	customerRepository repository.CustomerRepository = repository.NewCustomerRepository(db)
	jwtService         service.JwtService            = service.NewJwtService()
	authService        service.AuthService           = service.NewAuthService(customerRepository)
	customerService    service.CustomerService       = service.NewCustomerService(customerRepository)
	//productService     service.ProductService        = service.NewProductService(productRepository)
	authController     controller.AuthController     = controller.NewAuthController(authService, jwtService)
	customerController controller.CustomerController = controller.NewCustomerController(customerService, jwtService)
	//productController  controller.ProductController  = controller.NewProductController(productService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	//
	//authRoutes := r.Group("api/auth")
	//{
	//	authRoutes.POST("/login", authController.Login)
	//	authRoutes.POST("/register", authController.Register)
	//}
	//
	//userRoutes := r.Group("api/customer", middleware.AuthorizeJwt(jwtService))
	//{
	//	userRoutes.GET("/profile", customerController.GetProfile)
	//	userRoutes.PUT("/profile", customerController.UpdateProfile)
	//}

	//productRoutes := r.Group("api/product", middleware.AuthorizeJwt(jwtService))
	//{
	//	productRoutes.GET("/", productController.All)
	//	productRoutes.POST("/", productController.Insert)
	//	productRoutes.GET("/:id", productController.FindById)
	//	productRoutes.PUT("/:id", productController.Update)
	//	productRoutes.DELETE("/:id", productController.Delete)
	//}

	r.Run()
}
