package router

import (
	"go-gin-rest-api-with-jwt/controllers"
	"go-gin-rest-api-with-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		// productRouter.PUT("/:productId", middlewares.ProductAuthorization(), controllers.UpdateProduct)
	}

	return r
}
