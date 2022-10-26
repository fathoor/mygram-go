package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"not null" form:"title" valid:"required~Title is required"`
	Caption  string `json:"caption" form:"caption"`
	PhotoUrl string `json:"photo_url" gorm:"not null" form:"photo_url" valid:"required~Photo URL is required"`
	UserId   int    `json:"user_id" form:"user_id"`
	User     *User  `json:"user"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(p)

	if e != nil {
		return e
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(p)

	if e != nil {
		return e
	}

	return
}
