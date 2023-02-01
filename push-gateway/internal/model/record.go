package model

import (
	"push-gateway/internal/service"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Content    string
	Status     string
	Platform   string
	Error      string
	TemplateId uint
	UserId     uint
	HouseId    uint
	ContractId uint
	RecordId   uint
}

func (*Record) TableName() string {
	return "push_sms_record"
}

func (v *Record) Create(req *service.SmsCreateRequest) error {
	v.Content = req.Content
	v.Platform = "nowcn"
	v.Status = "init"
	v.UserId = uint(req.UserId)
	v.HouseId = uint(req.HouseId)
	v.ContractId = uint(req.ContractId)
	v.RecordId = uint(req.RecordId)
	v.TemplateId = 0

	if err := DB.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func BuildRecord(v *Record) *service.SmsRecordModel {
	return &service.SmsRecordModel{
		Id:         uint32(v.ID),
		UserId:     uint32(v.UserId),
		Content:    v.Content,
		Status:     v.Status,
		Error:      v.Error,
		CreatedAt:  v.CreatedAt.Format("2006-01-02 15:04:05"),
		HouseId:    uint32(v.HouseId),
		ContractId: uint32(v.ContractId),
		RecordId:   uint32(v.RecordId),
	}
}
