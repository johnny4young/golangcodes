package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string    `json:"username" gorm:"not null;unique"`
	Email          string    `json:"email" gorm:"not null;unique"`
	Fullname       string    `json:"fullname" gorm:"not null"`
	Password       string    `json:"password,omitempty" gorm:"not null;type:varchar(256)" `
	ConfirPassword string    `json:"confirmPassword,omitempty" gorm:"-"`
	Picture        string    `json:"picture"`
	Comments       []Comment `json:"comments,omitempty"`
}
