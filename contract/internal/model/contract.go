package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	UserId       uint
	HouseId      uint
	Status       string
	StartAt      time.Time
	EndAt        time.Time
	UntilAt      time.Time
	Money        float32
	Cycle        uint32
	TargetName   string
	TargetPhone  string
	TargetWxcode string
	Period       uint
	PeriodUnit   string
	Remark       string
	Deposit      float32
	OwnerMoney   float32
	OwnerUserId  uint
}
