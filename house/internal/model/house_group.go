package model

import (
	"house/internal/service"

	"gorm.io/gorm"
)

type HouseGroup struct {
	gorm.Model
	Name   string
	Status string
	UserId uint
}

func (group *HouseGroup) CheckExistById(id uint, user_id uint) bool {
	if err := DB.Where("id=? and user_id=?", id, user_id).First(&group).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (group *HouseGroup) Join(house_id uint) error {
	var rela HouseGroupRelation
	err := DB.FirstOrCreate(&rela, HouseGroupRelation{
		GroupId: group.ID,
		HouseId: house_id,
	}).Error
	return err
}
func (group *HouseGroup) Leave(house_id uint) error {
	var rela HouseGroupRelation
	err := DB.Where("group_id=? and house_id=?", group.ID, house_id).Unscoped().Delete(&rela).Error
	return err
}

func BuildHouseGroup(v *HouseGroup) *service.HouseGroupModel {
	return &service.HouseGroupModel{
		Id:        uint32(v.ID),
		Name:      v.Name,
		Status:    v.Status,
		UserId:    uint32(v.UserId),
		CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: v.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
