package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name           string `json:"name" gorm:"not null" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `json:"social_media_url" gorm:"not null" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserId         int    `json:"user_id" form:"user_id"`
	User           *User  `json:"user"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(s)

	if e != nil {
		return e
	}

	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, e := govalidator.ValidateStruct(s)

	if e != nil {
		return e
	}

	return
}
