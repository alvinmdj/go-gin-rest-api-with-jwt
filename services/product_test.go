package services

import (
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductSvc_FindAllSvcNotFound(t *testing.T) {
	productRepo := &repositories.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productRepo.Mock.On("FindAll").Return([]models.Product{}, "product not found")

	product, err := productSvcMock.FindAllSvc()

	assert.Equal(t, product, []models.Product{}, "product is an empty slice of struct")
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductSvc_FindAllSvcFound(t *testing.T) {
	productRepo := &repositories.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productsData := []models.Product{
		{
			Title:       "product 2",
			Description: "description of product 2",
			UserID:      2,
		},
		{
			Title:       "product 3",
			Description: "description of product 3",
			UserID:      2,
		},
	}

	productRepo.Mock.On("FindAll").Return(productsData, nil)

	products, err := productSvcMock.FindAllSvc()

	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, productsData[0].Title, products[0].Title, "product found is 'product 2'")
	assert.Equal(t, productsData[1].Title, products[1].Title, "product found is 'product 3'")
}

func TestProductSvc_FindByIdSvcNotFound(t *testing.T) {
	productRepo := &repositories.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productRepo.Mock.On("FindById", 1).Return(models.Product{}, "product not found")

	product, err := productSvcMock.FindByIdSvc(1)

	assert.Equal(t, product, models.Product{}, "product is an empty struct")
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductSvc_FindByIdSvcFound(t *testing.T) {
	productRepo := &repositories.ProductRepoMock{Mock: mock.Mock{}}
	productSvcMock := productSvc{productRepo: productRepo}

	productOne := models.Product{
		Title:       "product 2",
		Description: "description of product 2",
		UserID:      2,
	}

	productRepo.Mock.On("FindById", 2).Return(productOne, nil)

	product, err := productSvcMock.FindByIdSvc(2)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, productOne.Title, product.Title, "product found is 'product 2'")
	assert.Equal(t, productOne, product, "product found with id '2'")
}
