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
