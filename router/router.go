package router

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/handlers"
	"go-gin-rest-api-with-jwt/middlewares"
	"go-gin-rest-api-with-jwt/repositories"
	"go-gin-rest-api-with-jwt/services"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db := database.GetDB()

	userRepo := repositories.UserRepoImpl(db)
	userSvc := services.UserSvcImpl(userRepo)
	userHdl := handlers.UserHdlImpl(userSvc)

	productRepo := repositories.ProductRepoImpl(db)
	productSvc := services.ProductSvcImpl(productRepo)
	productHdl := handlers.ProductHdlImpl(productSvc)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		userRouter := v1.Group("/users")
		{
			userRouter.POST("/register", userHdl.Register)
			userRouter.POST("/login", userHdl.Login)
		}

		productRouter := v1.Group("/products")
		{
			productRouter.Use(middlewares.Authentication())
			productRouter.POST("/", productHdl.CreateHdl)

			productAuthorizedRouter := productRouter.Group("/")
			{
				productAuthorizedRouter.Use(middlewares.ProductAuthorization())
				productAuthorizedRouter.GET("/", productHdl.FindAllHdl)
				productAuthorizedRouter.GET("/:productId", productHdl.FindByIdHdl)
				productAuthorizedRouter.PUT("/:productId", productHdl.UpdateHdl)
				productAuthorizedRouter.DELETE("/:productId", productHdl.DeleteHdl)
			}
		}
	}

	return r
}
