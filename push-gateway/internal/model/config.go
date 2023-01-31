package model

import (
	"errors"

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
