package models

import (
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/farrelahmady/my-gram-hacktiv8/helpers"
	"gorm.io/gorm"
)

type User struct {
	Model
	Username string `gorm:"not null;unique" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;unique" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Your password must be at least 6 characters"`
	Age      uint   `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,type(uint)~Your age must be a number,range(8|100)~Your age must be between 1 and 100"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)
	u.Age = uint(u.Age)
	err = nil
	return

}
