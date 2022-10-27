package model

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/fathoor/mygram-go/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" gorm:"unique;not null" form:"email" valid:"email~Email must be valid,required~Email is required"`
	Password string `json:"password" gorm:"not null" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Age      int    `json:"age" gorm:"not null" form:"age" valid:"required~Age is required,range(8|100)~You must be at least 8 years old"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(u)

	if e != nil {
		return e
	}

	u.Password = helper.HashPassword(u.Password)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Email == "" {
		err = errors.New("email is required")
		return
	} else if u.Username == "" {
		err = errors.New("username is required")
		return
	} else if u.Email == "" && u.Username == "" {
		err = errors.New("email and password is required")
		return
	} else if !govalidator.IsEmail(u.Email) {
		err = errors.New("email is not valid")
		return
	}

	return
}
