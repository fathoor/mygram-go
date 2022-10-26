package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Message string `json:"message" gorm:"not null" form:"message" valid:"required~Message is required"`
	UserId  int    `json:"user_id" form:"user_id"`
	User    *User  `json:"user"`
	PhotoId int    `json:"photo_id" form:"photo_id"`
	Photo   *Photo `json:"photo"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(c)

	if e != nil {
		return e
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(c)

	if e != nil {
		return e
	}

	return
}
