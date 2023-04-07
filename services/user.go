package services

import (
	"go-gin-rest-api-with-jwt/models"
	"go-gin-rest-api-with-jwt/repositories"
	"log"
)

type UserSvcInterface interface {
	CreateSvc(user models.User) (models.User, error)
	FindByEmailSvc(user models.User) (models.User, error)
}

type userSvc struct {
	userRepo repositories.UserRepoInterface
}

func UserSvcImpl(userRepo repositories.UserRepoInterface) UserSvcInterface {
	return &userSvc{
		userRepo: userRepo,
	}
}

func (u *userSvc) CreateSvc(user models.User) (models.User, error) {
	log.Printf("[INFO] %T CreateSvc invoked\n", u)
	createdUser, err := u.userRepo.Create(user)
	if err != nil {
		log.Printf("[ERROR] %T CreateSvc :%v\n", u, err)
	}
	return createdUser, err
}

func (u *userSvc) FindByEmailSvc(user models.User) (models.User, error) {
	log.Printf("[INFO] %T FindByEmailSvc invoked\n", u)
	foundUser, err := u.userRepo.FindByEmail(user)
	if err != nil {
		log.Printf("[ERROR] %T FindByEmailSvc :%v\n", u, err)
	}
	return foundUser, err
}
