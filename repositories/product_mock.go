package repositories

import (
	"errors"
	"go-gin-rest-api-with-jwt/models"

	"github.com/stretchr/testify/mock"
)

type ProductRepoMock struct {
	Mock mock.Mock
}

func (p *ProductRepoMock) FindAll() (products []models.Product, err error) {
	args := p.Mock.Called()
	if args.Get(1) != nil /* error */ {
		return []models.Product{}, errors.New("product not found")
	}

	products = args.Get(0).([]models.Product)
	return products, nil
}

func (p *ProductRepoMock) FindAllByUserId(userId uint) (products []models.Product, err error) {
	return
}

func (p *ProductRepoMock) FindById(productId int) (product models.Product, err error) {
	args := p.Mock.Called(productId)
	if args.Get(1) != nil /* error */ {
		return product, errors.New("product not found")
	}

	product = args.Get(0).(models.Product)
	return product, nil
}

func (p *ProductRepoMock) Create(product models.Product) (models.Product, error) {
	return models.Product{}, nil
}

func (p *ProductRepoMock) Update(product models.Product, productId int) (models.Product, error) {
	return models.Product{}, nil
}

func (p *ProductRepoMock) Delete(product models.Product) (err error) {
	return
}
