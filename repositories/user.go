package repositories

import (
	"go-gin-rest-api-with-jwt/models"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	Create(user models.User) (models.User, error)
	FindByEmail(user models.User) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func UserRepoImpl(db *gorm.DB) UserRepoInterface {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(user models.User) (models.User, error) {
	err := u.db.Debug().Create(&user).Error
	return user, err
}

func (u *userRepo) FindByEmail(user models.User) (models.User, error) {
	err := u.db.Debug().Where("email = ?", user.Email).Take(&user).Error
	return user, err
}
