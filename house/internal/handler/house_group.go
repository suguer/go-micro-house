package handler

import (
	"context"
	"house/internal/service"
)

type HouseGroupService struct {
	service.UnimplementedHouseGroupServiceServer
}

func NewHouseGroupService() *HouseGroupService {
	return &HouseGroupService{}
}

func (*HouseGroupService) Index(ctx context.Context, req *service.HouseGroupRequest) (resp *service.HouseGroupResponse, err error) {
	resp = new(service.HouseGroupResponse)
	resp.Code = 440
	return resp, nil
}
