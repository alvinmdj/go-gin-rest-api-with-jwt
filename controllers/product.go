package controllers

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/helpers"
	"go-gin-rest-api-with-jwt/models"
	"net/http"

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

	c.JSON(http.StatusCreated, gin.H{
		"id":          Product.ID,
		"title":       Product.Title,
		"description": Product.Description,
		"user_id":     Product.UserID,
	})
}
