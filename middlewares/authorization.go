package middlewares

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		// get route param "productId"
		param := c.Param("productId")
		productId, err := strconv.Atoi(param)
		if param != "" && err != nil {
			log.Println("ERROR =>", err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "invalid parameter",
			})
			return
		}

		// get token claims set from authentication middleware
		userData := c.MustGet("userData").(jwt.MapClaims)

		// get user id from token claims
		userId := uint(userData["id"].(float64))

		User := models.User{}
		Product := models.Product{}

		// get user role
		err = db.Debug().Select("role").First(&User, userId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		// user role cannot update or delete product, abort request
		if User.Role == "user" && (c.Request.Method == "PUT" || c.Request.Method == "DELETE") {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		// set user role
		c.Set("userRole", User.Role)

		// call next early if request get all products
		if c.Request.Method == "GET" && param == "" {
			c.Next()
			return
		}

		// get user_id column from products table with the product id
		err = db.Debug().Select("user_id").First(&Product, productId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "data not found",
				"message": "data doesn't exist",
			})
			return
		}

		if User.Role == "user" {
			// check if user id from db == user id from token claims
			if Product.UserID != userId {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
					"error":   "unauthorized",
					"message": "you are not allowed to access this data",
				})
				return
			}
		}

		c.Next()
	}
}
