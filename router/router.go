package router

import (
	"go-gin-rest-api-with-jwt/controllers"
	"go-gin-rest-api-with-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userRouter := v1.Group("/users")
		{
			userRouter.POST("/register", controllers.UserRegister)
			userRouter.POST("/login", controllers.UserLogin)
		}

		productRouter := v1.Group("/products")
		{
			productRouter.Use(middlewares.Authentication())
			productRouter.POST("/", controllers.CreateProduct)

			productAuthorizedRouter := productRouter.Group("/")
			{
				productAuthorizedRouter.Use(middlewares.ProductAuthorization())
				productAuthorizedRouter.GET("/", controllers.GetAllProducts)
				productAuthorizedRouter.GET("/:productId", controllers.GetProductById)
				productAuthorizedRouter.PUT("/:productId", controllers.UpdateProduct)
				productAuthorizedRouter.DELETE("/:productId", controllers.DeleteProduct)
			}
		}
	}

	return r
}
