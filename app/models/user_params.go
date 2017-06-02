package models

import (
	"github.com/jinzhu/gorm"
)

type UserParam struct {
	gorm.Model
	UserID  uint
	Stamina uint
}

func NewUserParam() *UserParam {
	return new(UserParam)
}
