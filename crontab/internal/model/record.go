package model

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	HouseID      int64      `gorm:"column:house_id;type:int(11);not null" json:"house_id"`
	ContractID   int64      `gorm:"column:contract_id;type:int(11);not null" json:"contract_id"`
	Status       string     `gorm:"column:status;type:varchar(16);not null;default:normal" json:"status"`
	Money        float64    `gorm:"column:money;type:double(10,2);not null;default:0.00" json:"money"`
	Remark       *string    `gorm:"column:remark;type:varchar(128)" json:"remark"`
	AcceptAt     *time.Time `gorm:"column:accept_at;type:date" json:"accept_at"`
	StartAt      *time.Time `gorm:"column:start_at;type:date" json:"start_at"`
	EndAt        *time.Time `gorm:"column:end_at;type:date" json:"end_at"`
	Type         string     `gorm:"column:type;type:varchar(16);not null;default:rent" json:"type"`
	UserID       int64      `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	CompleteAt   *time.Time `gorm:"column:complete_at;type:datetime" json:"complete_at"`
	AccountID    int64      `gorm:"column:account_id;type:int(11);not null" json:"account_id"`
	SubmitAt     *time.Time `gorm:"column:submit_at;type:datetime" json:"submit_at"`
	SubmitUserID int64      `gorm:"column:submit_user_id;type:int(11);not null" json:"submit_user_id"`
	Append       *string    `gorm:"column:append;type:varchar(2048)" json:"append"`
	Desc         *string    `gorm:"column:desc;type:varchar(64)" json:"desc"`
	GroupID      int64      `gorm:"column:group_id;type:int(11);not null" json:"group_id"`
	ActualMoney  float64    `gorm:"column:actual_money;type:double(10,2);not null;default:0.00" json:"actual_money"`
}

func (*Record) TableName() string {
	return "cloud_rent_record"
}
