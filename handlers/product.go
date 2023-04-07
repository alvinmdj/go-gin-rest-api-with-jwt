package handlers

import (
	"fmt"
	"go-gin-rest-api-with-jwt/helpers"
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var appJSON = "application/json"

type ProductHdlInterface interface {
	FindAllHdl(c *gin.Context)
	FindByIdHdl(c *gin.Context)
	CreateHdl(c *gin.Context)
	UpdateHdl(c *gin.Context)
	DeleteHdl(c *gin.Context)
}

type productHdl struct {
	productService services.ProductSvcInterface
}

func ProductHdlImpl(productService services.ProductSvcInterface) ProductHdlInterface {
	return &productHdl{
		productService: productService,
	}
}

func (p *productHdl) FindAllHdl(c *gin.Context) {
	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	// get user role set from authorization middleware
	userRole := c.MustGet("userRole")

	if userRole == "user" {
		// only get user created products
		products, err := p.productService.FindAllByUserIdSvc(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	} else {
		// admin can access all data
		products, err := p.productService.FindAllSvc()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}
}

func (p *productHdl) FindByIdHdl(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))
	product, err := p.productService.FindByIdSvc(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (p *productHdl) CreateHdl(c *gin.Context) {
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	// get token claims in userData context from authentication middleware
	// and cast the data type from any to jwt.MapClaims
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	// set user id from token claims into the Product struct
	product.UserID = userId

	createdProduct, err := p.productService.CreateSvc(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (p *productHdl) UpdateHdl(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userId
	product.ID = uint(productId)

	updatedProduct, err := p.productService.UpdateSvc(product, productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (p *productHdl) DeleteHdl(c *gin.Context) {
	product := models.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	product.ID = uint(productId)

	if err := p.productService.DeleteSvc(product); err != nil {
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
