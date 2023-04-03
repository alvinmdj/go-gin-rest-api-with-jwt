package models

import (
	"go-gin-rest-api-with-jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// user role: user | admin

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     string    `gorm:"not null;default:user" json:"role"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// validate input
	if _, err = govalidator.ValidateStruct(u); err != nil {
		return
	}

	// hash password
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return
}
