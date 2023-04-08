package services

import (
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepo = &repositories.ProductRepoMock{Mock: mock.Mock{}}
var productSvcMock = productSvc{productRepo: productRepo}

func TestProductSvc_FindByIdNotSvcFound(t *testing.T) {
	productRepo.Mock.On("FindById", 1).Return(models.Product{}, "product not found")

	product, err := productSvcMock.FindByIdSvc(1)

	assert.Equal(t, product, models.Product{}, "product is an empty struct")
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductSvc_FindByIdSvcFound(t *testing.T) {
	productOne := models.Product{
		Title:       "product 2",
		Description: "description of product 2",
		UserID:      2,
	}

	productRepo.Mock.On("FindById", 2).Return(productOne, nil)

	product, err := productSvcMock.FindByIdSvc(2)
	_ = product

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productOne.Title, product.Title, "product found is 'product 2'")
	assert.Equal(t, productOne, product, "product found with id '2'")
}
