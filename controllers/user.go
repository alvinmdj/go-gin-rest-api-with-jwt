package controllers

import (
	"go-gin-rest-api-with-jwt/database"
	"go-gin-rest-api-with-jwt/helpers"
	"go-gin-rest-api-with-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	if err := db.Debug().Create(&User).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": gin.H{
			"id":        User.ID,
			"full_name": User.FullName,
			"email":     User.Email,
		},
	})
}
