package services

import (
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/repositories"
	"log"
)

type ProductSvcInterface interface {
	FindAllSvc() (products []models.Product, err error)
	FindAllByUserIdSvc(userId uint) (products []models.Product, err error)
	FindByIdSvc(productId int) (product models.Product, err error)
	CreateSvc(product models.Product) (models.Product, error)
	UpdateSvc(product models.Product, productId int) (models.Product, error)
	DeleteSvc(product models.Product) (err error)
}

type productSvc struct {
	productRepo repositories.ProductRepoInterface
}

func ProductSvcImpl(productRepo repositories.ProductRepoInterface) ProductSvcInterface {
	return &productSvc{
		productRepo: productRepo,
	}
}

func (p *productSvc) FindAllSvc() (products []models.Product, err error) {
	log.Printf("[INFO] %T FindAllSvc invoked\n", p)
	if products, err = p.productRepo.FindAll(); err != nil {
		log.Printf("[ERROR] %T FindAllSvc :%v\n", p, err)
	}
	return
}

func (p *productSvc) FindAllByUserIdSvc(userId uint) (products []models.Product, err error) {
	log.Printf("[INFO] %T FindAllByUserIdSvc invoked\n", p)
	if products, err = p.productRepo.FindAllByUserId(userId); err != nil {
		log.Printf("[ERROR] %T FindAllByUserIdSvc :%v\n", p, err)
	}
	return
}

func (p *productSvc) FindByIdSvc(productId int) (product models.Product, err error) {
	log.Printf("[INFO] %T FindByIdSvc invoked\n", p)
	if product, err = p.productRepo.FindById(productId); err != nil {
		log.Printf("[ERROR] %T FindByIdSvc :%v\n", p, err)
	}
	return
}

func (p *productSvc) CreateSvc(product models.Product) (models.Product, error) {
	log.Printf("[INFO] %T CreateSvc invoked\n", p)
	createdProduct, err := p.productRepo.Create(product)
	if err != nil {
		log.Printf("[ERROR] %T CreateSvc :%v\n", p, err)
	}
	return createdProduct, err
}

func (p *productSvc) UpdateSvc(product models.Product, productId int) (models.Product, error) {
	log.Printf("[INFO] %T UpdateSvc invoked\n", p)
	updatedProduct, err := p.productRepo.Update(product, productId)
	if err != nil {
		log.Printf("[ERROR] %T UpdateSvc :%v\n", p, err)
	}
	return updatedProduct, err
}

func (p *productSvc) DeleteSvc(product models.Product) (err error) {
	log.Printf("[INFO] %T DeleteSvc invoked\n", p)
	if err = p.productRepo.Delete(product); err != nil {
		log.Printf("[ERROR] %T DeleteSvc :%v\n", p, err)
	}
	return
}
