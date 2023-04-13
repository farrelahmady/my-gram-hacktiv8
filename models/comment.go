package models

type Comment struct {
	Model
	Message string `gorm:"type:varchar(255);not null" json:"message" form:"message" valid:"required~Your message is required"`
	UserID  uint
	User    *User
	PhotoID uint
	Photo   *Photo
}

func (c *Comment) TableName() string {
	return "comments"
}
