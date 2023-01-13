package model

import (
	"user/internal/service"

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
	DB.Where("user_id=?", UserId).Find(&list)
	return list, nil
}

func BuildUserPlatform(v UserPlatform) *service.UserPlatformModel {
	p := service.UserPlatformModel{
		Id:        uint32(v.ID),
		UserId:    uint32(v.UserId),
		Platform:  v.Platform,
		CreatedAt: v.CreatedAt.Format("2006-01-02"),
	}
	return &p
}
