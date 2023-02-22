package model

import (
	"time"

	"gorm.io/gorm"
)

// CloudRentHouse mapped from table <cloud_rent_house>
type House struct {
	ID                    int64          `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Title                 *string        `gorm:"column:title;type:varchar(64)" json:"title"`
	UserID                int64          `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	Cycle                 int64          `gorm:"column:cycle;type:int(11);not null;default:1" json:"cycle"`
	Area                  float64        `gorm:"column:area;type:double(10,2);not null;default:0.00" json:"area"`
	ContractID            int64          `gorm:"column:contract_id;type:int(11);not null" json:"contract_id"`
	Status                string         `gorm:"column:status;type:varchar(16);not null;default:normal" json:"status"`
	Money                 float64        `gorm:"column:money;type:double(10,2);not null;default:0.00" json:"money"`
	CreatedAt             *time.Time     `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt             *time.Time     `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp" json:"-"`
	Lat                   float64        `gorm:"column:lat;type:double(10,6);not null;default:0.000000" json:"lat"`
	Lng                   float64        `gorm:"column:lng;type:double(10,6);not null;default:0.000000" json:"lng"`
	Estate                *string        `gorm:"column:estate;type:varchar(32)" json:"estate"`
	Address               *string        `gorm:"column:address;type:varchar(256)" json:"address"`
	Province              *string        `gorm:"column:province;type:varchar(16)" json:"province"`
	City                  *string        `gorm:"column:city;type:varchar(16)" json:"city"`
	Region                *string        `gorm:"column:region;type:varchar(16)" json:"region"`
	Linkname              *string        `gorm:"column:linkname;type:varchar(16)" json:"linkname"`
	Linkphone             *string        `gorm:"column:linkphone;type:varchar(16)" json:"linkphone"`
	Remark                *string        `gorm:"column:remark;type:varchar(256)" json:"remark"`
	Introduction          *string        `gorm:"column:introduction;type:varchar(512)" json:"introduction"`
	Pv                    int64          `gorm:"column:pv;type:int(11);not null" json:"pv"`
	Cv                    int64          `gorm:"column:cv;type:int(11);not null" json:"cv"`
	RoomConfig            *string        `gorm:"column:room_config;type:varchar(64)" json:"room_config"`
	Electricmeter         *string        `gorm:"column:electricmeter;type:varchar(32)" json:"electricmeter"`
	Watermeter            *string        `gorm:"column:watermeter;type:varchar(32)" json:"watermeter"`
	Storey                int64          `gorm:"column:storey;type:int(11);not null" json:"storey"`
	CommunityName         *string        `gorm:"column:community_name;type:varchar(32)" json:"community_name"`
	Hardware              *string        `gorm:"column:hardware;type:varchar(2048)" json:"hardware"`
	ElectricityFee        float64        `gorm:"column:electricity_fee;type:double(10,2);not null;default:0.00" json:"electricity_fee"`
	WaterFee              float64        `gorm:"column:water_fee;type:double(10,2);not null;default:0.00" json:"water_fee"`
	ManageFee             float64        `gorm:"column:manage_fee;type:double(10,2);not null;default:0.00" json:"manage_fee"`
	Fee                   *string        `gorm:"column:fee;type:varchar(512)" json:"fee"`
	ActiveContractCount   int64          `gorm:"-:all" json:"ActiveContractCount"`
	UnActiveContractCount int64          `gorm:"-:all" json:"UnActiveContractCount"`
	IsOwner               int64          `gorm:"-:all" json:"isOwner"`
}

// TableName CloudRentHouse's table name
func (*House) TableName() string {
	return "cloud_rent_house"
}
