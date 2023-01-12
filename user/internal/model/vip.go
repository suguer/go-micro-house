package model

import (
	"time"
	"user/internal/service"

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

func (vip *Vip) FindActiveById(UserId string) error {
	if err := DB.Where("user_id=? and status=?", UserId, "normal").First(&vip).Error; err == gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func BuildVip(item Vip) *service.UserVipModel {
	vip := service.UserVipModel{
		Id:        uint32(item.ID),
		UserId:    uint32(item.UserId),
		CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		StartAt:   item.StartAt.Format("2006-01-02"),
		EndAt:     item.EndAt.Format("2006-01-02"),
		Level:     item.Level,
		Status:    item.Status,
	}
	return &vip
}
