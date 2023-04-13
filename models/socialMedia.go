package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	Model
	Name           string `gorm:"type:varchar(255);not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaURL string `gorm:"type:varchar(255);not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social_media_url is required"`
	UserID         uint
	User           *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (s *SocialMedia) TableName() string {
	return "social_media"
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
