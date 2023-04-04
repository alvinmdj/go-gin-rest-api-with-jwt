package controllers

import (
	"fmt"
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/helpers"
	"go-gin-rest-api-with-jwt/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	// set user id from token claims into the Product struct
	Product.UserID = userId

	if err := db.Debug().Create(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func GetAllProducts(c *gin.Context) {
	db := database.GetDB()
	Products := []models.Product{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// get user role set from authorization middleware
	userRole := c.MustGet("userRole")

	if userRole == "user" {
		// only get user created products
		if err := db.Debug().Where("user_id = ?", userId).Find(&Products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": err.Error(),
			})
			return
		}
	} else {
		// admin can access all data
		if err := db.Debug().Find(&Products).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, Products)
}

func GetProductById(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))

	if err := db.Debug().First(&Product, productId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userId
	Product.ID = uint(productId)

	err := db.Debug().Model(&Product).
		Where("id = ?", productId).
		Updates(models.Product{
			Title:       Product.Title,
			Description: Product.Description,
		}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	Product.ID = uint(productId)

	if err := db.Debug().Delete(&Product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product with id %d has been deleted", productId),
	})
}
