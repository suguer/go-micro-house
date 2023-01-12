package model

import (
	"time"

	"gorm.io/gorm"
)

type Vip struct {
	gorm.Model
	UserId  uint
	StartAt time.Time
	EndAt   time.Time
	Level   string
	Status  string
}
