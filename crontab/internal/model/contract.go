package model

import (
	"time"

	"gorm.io/gorm"
)

// CloudRentContract mapped from table <cloud_rent_contract>
type Contract struct {
	ID            int64          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	HouseID       int64          `gorm:"column:house_id;type:int(11);not null" json:"house_id"`
	Status        *string        `gorm:"column:status;type:varchar(16);default:normal" json:"status"`
	StartAt       *time.Time     `gorm:"column:start_at;type:date" json:"start_at"`
	EndAt         *time.Time     `gorm:"column:end_at;type:date" json:"end_at"`
	UntilAt       time.Time      `gorm:"column:until_at;type:date;not null" json:"until_at"`
	Money         float64        `gorm:"column:money;type:double(10,2);not null;default:0.00" json:"money"`
	Cycle         int64          `gorm:"column:cycle;type:int(11);not null;default:1" json:"cycle"`
	TargetName    *string        `gorm:"column:target_name;type:varchar(32)" json:"target_name"`
	TargetPhone   *string        `gorm:"column:target_phone;type:varchar(16)" json:"target_phone"`
	TargetWxcode  *string        `gorm:"column:target_wxcode;type:varchar(32)" json:"target_wxcode"`
	CreatedAt     *time.Time     `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt     *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"-"`
	UserID        int64          `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	Period        int64          `gorm:"column:period;type:int(11);not null;default:1" json:"period"`
	PeriodUnit    string         `gorm:"column:period_unit;type:varchar(16);not null;default:month" json:"period_unit"`
	Remark        *string        `gorm:"column:remark;type:varchar(256)" json:"remark"`
	Deposit       float64        `gorm:"column:deposit;type:double(10,2);not null;default:0.00" json:"deposit"`
	WaterNo       float64        `gorm:"column:water_no;type:double(10,6);not null;default:0.000000" json:"water_no"`
	ElectricityNo float64        `gorm:"column:electricity_no;type:double(10,6);not null;default:0.000000" json:"electricity_no"`
	OwnerMoney    float64        `gorm:"column:owner_money;type:double(10,2);not null;default:0.00" json:"owner_money"`
	OwnerUserID   int64          `gorm:"column:owner_user_id;type:int(11);not null" json:"owner_user_id"`
	House         House          `gorm:"references:ID"`
}

// TableName CloudRentContract's table name
func (*Contract) TableName() string {
	return "cloud_rent_contract"
}
