package router

import (
	"go-gin-rest-api-with-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		// userRouter.POST("/login", controllers.UserLogin)
	}

	return r
}
