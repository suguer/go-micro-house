package model

import (
	"house/internal/service"

	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	UserId     uint
	ContractId uint
	Cycle      uint32
	Title      string
	Status     string
	Money      float32
	Province   string
	City       string
	Region     string
	Linkname   string
	Linkphone  string
	Remark     string
}

func (house *House) CheckExistById(id uint, user_id uint) bool {
	if err := DB.Where("id=? and user_id=?", id, user_id).First(&house).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
func (house *House) CheckExistByName(name string, user_id uint) bool {
	if err := DB.Where("name=? and user_id=?", name, user_id).First(&house).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (house *House) CheckCanCreate(level string) bool {
	arr := map[string]int{
		"":         5,
		"gold":     50,
		"bojin":    100,
		"diamonds": -1,
	}
	if value, ok := arr[level]; ok {
		if value == -1 {
			return true
		}
		var count int64
		DB.Where("user_id=?", house.UserId).Count(&count)
		if count < int64(value) {
			return true
		}
	}
	return false
}

func BuildHouse(house *House) *service.HouseModel {
	return &service.HouseModel{
		Id:         uint32(house.ID),
		UserId:     uint32(house.UserId),
		Title:      house.Title,
		Cycle:      house.Cycle,
		Status:     house.Status,
		ContractId: uint32(house.ContractId),
		Money:      house.Money,
		Province:   house.Province,
		City:       house.City,
		Region:     house.Region,
		Linkname:   house.Linkname,
		Linkphone:  house.Linkphone,
		Remark:     house.Remark,
	}
}
