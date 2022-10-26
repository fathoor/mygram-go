package model

import (
	"github.com/asaskevich/govalidator"
	"github.com/fathoor/mygram-go/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" form:"username" valid:"required~Username is required"`
	Email    string `json:"email" gorm:"unique;not null" form:"email" valid:"email~Email must be valid,required~Email is required"`
	Password string `json:"password" gorm:"not null" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Age      int    `json:"age" gorm:"not null" form:"age" valid:"required~Age is required,min(8)~Age must be at least 8"`
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
	_, e := govalidator.ValidateStruct(u)

	if e != nil {
		return e
	}

	u.Password = helper.HashPassword(u.Password)

	return
}
