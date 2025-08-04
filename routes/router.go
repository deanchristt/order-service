package routes

import (
	"github.com/deanchristt/order-service/controller"
	"github.com/deanchristt/order-service/middleware"
	"github.com/deanchristt/order-service/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	authController controller.AuthController,
	userController controller.UserController,
	jwtService service.JwtService,
	productController controller.ProductController,
) {
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/customer", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.GET("/profile", userController.GetProfile)
		userRoutes.PUT("/profile", userController.UpdateProfile)
	}

	productRoutes := r.Group("api/product", middleware.AuthorizeJwt(jwtService))
	{
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/", productController.Insert)
		productRoutes.GET("/:id", productController.FindById)
		productRoutes.PUT("/:id", productController.Update)
		productRoutes.DELETE("/:id", productController.Delete)
	}
}
