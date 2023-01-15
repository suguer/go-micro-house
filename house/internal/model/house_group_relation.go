package model

import "gorm.io/gorm"

type HouseGroupRelation struct {
	gorm.Model
	HouseId uint
	GroupId uint
}
