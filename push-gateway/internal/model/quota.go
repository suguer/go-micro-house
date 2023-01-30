package model

import (
	"errors"

	"gorm.io/gorm"
)

type Quota struct {
	gorm.Model
	UserId    uint
	Available uint
	Total     uint
}

func (*Quota) TableName() string {
	return "push_sms_quota"
}

func (v *Quota) CheckAvailableCount(user_id uint) (uint, error) {
	if err := DB.Where("user_id=?", user_id).First(&v).Error; err != nil {
		return 0, err
	}
	if v.Available == 0 {
		return v.Available, errors.New("limit")
	}
	return v.Available, nil
}

func (v *Quota) Consume(count uint) error {
	v.Available = v.Available - count
	DB.Save(&v)
	return nil
}
