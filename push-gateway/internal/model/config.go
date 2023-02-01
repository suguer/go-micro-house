package model

import (
	"errors"
	"push-gateway/internal/service"

	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	UserId     uint
	Available  uint
	Total      uint
	Status     string
	Day        string
	Expiration uint
}

func (*Config) TableName() string {
	return "push_sms_config"
}

func (v *Config) CheckAvailableCount(user_id uint) (uint, error) {
	if err := DB.Where("user_id=?", user_id).First(&v).Error; err != nil {
		return 0, err
	}
	if v.Available == 0 {
		return v.Available, errors.New("limit")
	}
	return v.Available, nil
}

func (v *Config) Consume(count uint) error {
	v.Available = v.Available - count
	DB.Save(&v)
	return nil
}

func (v *Config) Build() *service.SmsConfigModel {
	return &service.SmsConfigModel{
		UserId:     uint32(v.UserId),
		Available:  uint32(v.Available),
		Total:      uint32(v.Total),
		Status:     v.Status,
		Day:        v.Day,
		Expiration: uint32(v.Expiration),
		// CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
