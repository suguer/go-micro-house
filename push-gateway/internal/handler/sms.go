package handler

import (
	"context"
	"fmt"
	"push-gateway/internal/gateway"
	"push-gateway/internal/model"
	"push-gateway/internal/service"
)

type SmsService struct {
	service.UnimplementedSmsServiceServer
}

func NewSmsService() *SmsService {
	return &SmsService{}
}

func (*SmsService) Index(ctx context.Context, req *service.SmsIndexRequest) (resp *service.SmsIndexResponse, err error) {
	resp = new(service.SmsIndexResponse)
	fmt.Printf("req: %v\n", req)
	pagination := model.NewPagination(req.Current, req.PageSize)
	var list []model.Record
	err = model.DB.Model(&model.Record{}).
		Scopes(pagination.Paginate()).
		Where("user_id=?", req.UserId).
		Find(&list).Limit(-1).Offset(-1).
		Count(&pagination.Total).Error
	resp.Pagination = pagination.Build()
	resp.Data = make([]*service.SmsRecordModel, len(list))
	for i, hm := range list {
		resp.Data[i] = model.BuildRecord(&hm)
	}
	return resp, err
}

func (*SmsService) Create(ctx context.Context, req *service.SmsRequest) (resp *service.SmsResponse, err error) {
	resp = new(service.SmsResponse)
	var config model.Config
	_, err = config.CheckAvailableCount(uint(req.UserId))
	if err != nil {
		return resp, err
	}
	var record model.Record
	err = record.Create(req)
	if err == nil {
		client := gateway.NowcnGateway{}
		if err = client.SendSms(); err != nil {
			record.Status = "fail"
			record.Error = err.Error()
			model.DB.Save(&record)
		} else {
			config.Consume(1)
			record.Status = "success"
			model.DB.Save(&record)
		}
	}
	return resp, err
}
