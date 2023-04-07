package handlers

import (
	"go-gin-rest-api-with-jwt/helpers"
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHdlInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userHdl struct {
	userSvc services.UserSvcInterface
}

func UserHdlImpl(userSvc services.UserSvcInterface) UserHdlInterface {
	return &userHdl{
		userSvc: userSvc,
	}
}

func (u *userHdl) Register(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	createdUser, err := u.userSvc.CreateSvc(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        createdUser.ID,
		"full_name": createdUser.FullName,
		"email":     createdUser.Email,
	})
}

func (u *userHdl) Login(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	userInput := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&userInput)
	} else {
		c.ShouldBind(&userInput)
	}

	password = userInput.Password

	// check if email is registered
	foundUser, err := u.userSvc.FindByEmailSvc(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	// compare password
	if isEqual := helpers.CompareHash([]byte(foundUser.Password), []byte(password)); !isEqual {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email or password",
		})
		return
	}

	// generate jwt
	token := helpers.GenerateToken(foundUser.ID, foundUser.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
