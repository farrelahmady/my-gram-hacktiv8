package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Model
	Title    string    `gorm:"type:varchar(255);not null" json:"title" form:"title" valid:"required~Your title is required"`
	Caption  string    `gorm:"type:varchar(255)" json:"caption" form:"caption"`
	PhotoURL string    `gorm:"type:varchar(255);not null" json:"photo_url" form:"photo_url" valid:"required~Your photo_url is required"`
	Comments []Comment `gorm:constraint:OnUpdate:CASCADE,OnDelete:SET NULL; json:"comments"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
