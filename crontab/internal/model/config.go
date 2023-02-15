package model

import (
	"errors"

	"gorm.io/gorm"
)

type SmsConfig struct {
	gorm.Model
	UserId     uint
	Available  uint
	Total      uint
	Status     string
	Day        string
	Expiration uint
}

func (*SmsConfig) TableName() string {
	return "push_sms_config"
}

func (v *SmsConfig) CheckAvailableCount(user_id uint) (uint, error) {
	if err := DB.Where("user_id=?", user_id).First(&v).Error; err != nil {
		return 0, err
	}
	if v.Available == 0 {
		return v.Available, errors.New("limit")
	}
	return v.Available, nil
}

func (v *SmsConfig) Consume(count uint) error {
	v.Available = v.Available - count
	DB.Save(&v)
	return nil
}
