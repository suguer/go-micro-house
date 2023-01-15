package handler

import (
	"context"
	"fmt"
	"house/config"
	"house/internal/model"
	"house/internal/service"
	"testing"
	"time"
)

func TestCreateHouse(t *testing.T) {
	config.InitConfig()
	model.InitDB()
	var ctx context.Context
	ser := NewHouseService()
	resp, err := ser.Create(ctx, &service.HouseRequest{
		Data: &service.HouseModel{
			UserId: 1,
			Title:  "20230115",
		},
	})
	if resp.Code != 0 {
		t.Fatal(resp)
	}
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %v\n", resp)
}

func TestUpdateHouse(t *testing.T) {
	config.InitConfig()
	model.InitDB()
	var ctx context.Context
	ser := NewHouseService()
	resp, err := ser.Update(ctx, &service.HouseRequest{
		Data: &service.HouseModel{
			Id:     1,
			UserId: 1,
			Title:  time.Now().Format("2006-01-02 15:04:05"),
		},
	})
	if resp.Code != 0 {
		t.Fatal(resp.Code)
	}
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %v\n", resp)
}

func TestHouseGroup(t *testing.T) {
	config.InitConfig()
	model.InitDB()
	var ctx context.Context
	ser := NewHouseGroupService()
	resp, err := ser.Create(ctx, &service.HouseGroupRequest{
		Name:   time.Now().Format("2006-01-02 15:04:05"),
		UserId: 1,
	})
	if err != nil {
		t.Log(err)
	}
	ser.Join(ctx, &service.HouseGroupRequest{
		GroupId: resp.HouseGroupModel.Id,
		HouseId: 1,
	})
	if err != nil {
		t.Log(err)
	}
	ser.Leave(ctx, &service.HouseGroupRequest{
		GroupId: resp.HouseGroupModel.Id,
		HouseId: 1,
	})
	if err != nil {
		t.Log(err)
	}
}
