package handler

import (
	"context"
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
	var pagination service.Pagination
	list := make([]model.Record, req.PageSize)
	err = model.DB.Model(&model.Record{}).
		Scopes(model.Paginate(int(req.PageSize), int(req.Current))).
		Where("user_id=?", req.UserId).
		Find(&list).Limit(-1).Offset(-1).
		Count(&pagination.Total).Error
	resp.Pagination = &pagination
	resp.Data = make([]*service.SmsRecordModel, len(list))
	for i, hm := range list {
		resp.Data[i] = model.BuildRecord(&hm)
	}
	return resp, err
}

func (*SmsService) Create(ctx context.Context, req *service.SmsRequest) (resp *service.SmsResponse, err error) {
	resp = new(service.SmsResponse)
	var quota model.Quota
	_, err = quota.CheckAvailableCount(uint(req.UserId))
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
			quota.Consume(1)
			record.Status = "success"
			model.DB.Save(&record)
		}
	}
	return resp, err
}
