package handler

import (
	"context"
	"fmt"
	"house/internal/model"
	"house/internal/service"
)

type HouseGroupService struct {
	service.UnimplementedHouseGroupServiceServer
}

func NewHouseGroupService() *HouseGroupService {
	return &HouseGroupService{}
}

func (*HouseGroupService) Index(ctx context.Context, req *service.HouseGroupIndexRequest) (resp *service.HouseGroupIndexResponse, err error) {
	resp = new(service.HouseGroupIndexResponse)
	var list []model.HouseGroup
	var pagination service.Pagination
	err = model.DB.Scopes(model.Paginate(int(req.PageSize), int(req.Current))).
		Where("user_id=?", req.UserId).
		Find(list).Limit(-1).Offset(-1).
		Count(&pagination.Total).Error
	return resp, err
}
func (*HouseGroupService) Create(ctx context.Context, req *service.HouseGroupRequest) (resp *service.HouseGroupResponse, err error) {
	fmt.Printf("req.Name: %v\n", req.Name)
	resp = new(service.HouseGroupResponse)
	var group model.HouseGroup
	model.DB.FirstOrCreate(&group, model.HouseGroup{
		Name:   req.Name,
		UserId: uint(req.UserId),
	})
	resp.HouseGroupModel = model.BuildHouseGroup(&group)
	return resp, err
}

func (*HouseGroupService) Remove(ctx context.Context, req *service.HouseGroupRequest) (resp *service.HouseGroupResponse, err error) {
	resp = new(service.HouseGroupResponse)
	var group model.HouseGroup
	model.DB.Create(&group)
	resp.HouseGroupModel = model.BuildHouseGroup(&group)
	return resp, err
}

func (*HouseGroupService) Join(ctx context.Context, req *service.HouseGroupRequest) (resp *service.HouseGroupResponse, err error) {
	resp = new(service.HouseGroupResponse)
	var group model.HouseGroup
	group.ID = uint(req.GroupId)
	err = group.Join(uint(req.HouseId))
	return resp, err
}
func (*HouseGroupService) Leave(ctx context.Context, req *service.HouseGroupRequest) (resp *service.HouseGroupResponse, err error) {
	resp = new(service.HouseGroupResponse)
	var group model.HouseGroup
	group.ID = uint(req.GroupId)
	err = group.Leave(uint(req.HouseId))
	return resp, err
}
