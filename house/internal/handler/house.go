package handler

import (
	"context"
	"errors"
	"house/internal/model"
	"house/internal/service"
)

type HouseService struct {
	service.UnimplementedHouseServiceServer
}

func NewHouseService() *HouseService {
	return &HouseService{}
}

func (*HouseService) Index(ctx context.Context, req *service.HouseIndexRequest) (resp *service.HouseIndexResponse, err error) {
	resp = new(service.HouseIndexResponse)
	list := make([]model.House, 0)
	var pagination service.Pagination
	err = model.DB.Model(&model.House{}).
		Scopes(model.Paginate(int(req.PageSize), int(req.Current))).
		Where("user_id=?", req.UserId).
		Find(&list).Limit(-1).Offset(-1).
		Count(&pagination.Total).Error
	resp.Pagination = &pagination
	resp.Data = make([]*service.HouseModel, len(list))
	for i, hm := range list {
		resp.Data[i] = model.BuildHouse(&hm)
	}
	return resp, err
}

func (*HouseService) Instance(ctx context.Context, req *service.HouseRequest) (resp *service.HouseDetailResponse, err error) {
	resp = new(service.HouseDetailResponse)
	var house model.House
	exist := house.CheckExistById(uint(req.Data.Id), uint(req.Data.UserId))
	if !exist {
		resp.Code = 404
		return resp, errors.New("empty")
	}
	resp.Data = model.BuildHouse(&house)
	return resp, err
}

func (*HouseService) Create(ctx context.Context, req *service.HouseRequest) (resp *service.HouseDetailResponse, err error) {
	resp = new(service.HouseDetailResponse)
	var house model.House

	house.UserId = uint(req.Data.UserId)
	check := house.CheckCanCreate(req.Level)
	if !check {
		resp.Code = 401
		return resp, errors.New("forbidden")
	}
	house.Title = req.Data.Title
	house.City = req.Data.City
	house.Province = req.Data.Province
	house.Region = req.Data.Region
	house.Cycle = req.Data.Cycle
	house.Status = req.Data.Status
	house.Money = req.Data.Money
	house.Linkname = req.Data.Linkname
	house.Linkphone = req.Data.Linkphone
	house.Remark = req.Data.Remark
	house.ContractId = uint(req.Data.ContractId)
	err = model.DB.Create(&house).Error
	resp.Data = model.BuildHouse(&house)
	return resp, err
}

func (*HouseService) Update(ctx context.Context, req *service.HouseRequest) (resp *service.HouseDetailResponse, err error) {
	resp = new(service.HouseDetailResponse)
	var house model.House
	exist := house.CheckExistById(uint(req.Data.Id), uint(req.Data.UserId))
	if !exist {
		resp.Code = 404
		return resp, errors.New("empty")
	}
	house.Title = req.Data.Title
	house.City = req.Data.City
	house.Province = req.Data.Province
	house.Region = req.Data.Region
	house.Cycle = req.Data.Cycle
	house.Status = req.Data.Status
	house.Money = req.Data.Money
	house.Linkname = req.Data.Linkname
	house.Linkphone = req.Data.Linkphone
	house.Remark = req.Data.Remark
	house.ContractId = uint(req.Data.ContractId)
	err = model.DB.Save(&house).Error
	return resp, err
}

func (*HouseService) Remove(ctx context.Context, req *service.HouseRequest) (resp *service.HouseDetailResponse, err error) {
	resp = new(service.HouseDetailResponse)
	var house model.House
	exist := house.CheckExistById(uint(req.Data.Id), uint(req.Data.UserId))
	if !exist {
		resp.Code = 404
		return resp, errors.New("empty")
	}
	err = model.DB.Delete(&house).Error
	return resp, err
}
