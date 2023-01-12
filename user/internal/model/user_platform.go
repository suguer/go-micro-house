package model

import (
	"gorm.io/gorm"
)

type UserPlatform struct {
	gorm.Model
	Openid   string
	Token    string
	Platform string
	UserId   uint
}

func (platform *UserPlatform) SelectByUserId(UserId string) ([]UserPlatform, error) {
	var list []UserPlatform
	return list, nil
}
