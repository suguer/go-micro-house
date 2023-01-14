package handler

import (
	"context"
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
	return resp, nil
}
