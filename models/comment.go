package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Model
	Message string `gorm:"type:varchar(255);not null" json:"message" form:"message" valid:"required~Your message is required"`
	UserID  uint
	User    *User
	PhotoID uint `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Your photo_id is required"`
	Photo   *Photo
}

func (c *Comment) TableName() string {
	return "comments"
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
