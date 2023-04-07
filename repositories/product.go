package repositories

import (
	"go-gin-rest-api-with-jwt/models"

	"gorm.io/gorm"
)

type ProductRepoInterface interface {
	FindAll() (products []models.Product, err error)
	FindAllByUserId(userId uint) (products []models.Product, err error)
	FindById(productId int) (product models.Product, err error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product, productId int) (models.Product, error)
	Delete(product models.Product) (err error)
}

type productRepo struct {
	db *gorm.DB
}

func ProductRepoImpl(db *gorm.DB) ProductRepoInterface {
	return &productRepo{
		db: db,
	}
}

func (p *productRepo) FindAll() (products []models.Product, err error) {
	err = p.db.Debug().Find(&products).Error
	return
}

func (p *productRepo) FindAllByUserId(userId uint) (products []models.Product, err error) {
	err = p.db.Debug().Where("user_id = ?", userId).Find(&products).Error
	return
}

func (p *productRepo) FindById(productId int) (product models.Product, err error) {
	err = p.db.Debug().First(&product, productId).Error
	return
}

func (p *productRepo) Create(product models.Product) (models.Product, error) {
	err := p.db.Debug().Create(&product).Error
	return product, err
}

func (p *productRepo) Update(product models.Product, productId int) (models.Product, error) {
	err := p.db.Debug().Model(&product).
		Where("id = ?", productId).
		Updates(models.Product{
			Title:       product.Title,
			Description: product.Description,
		}).Error
	return product, err
}

func (p *productRepo) Delete(product models.Product) (err error) {
	err = p.db.Debug().Delete(&product).Error
	return
}
