package handler

import (
	"context"
	"user/internal/model"
	"user/internal/service"
)

type UserService struct {
	service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) Login(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user model.User
	resp = new(service.UserDetailResponse)
	err = user.ShowUserInfo(req)
	if err != nil {
		return resp, err
	}
	resp.UserDetail = model.BuildUser(user)
	return resp, nil
}
func (*UserService) Info(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user model.User
	resp = new(service.UserDetailResponse)
	model.DB.Model(&model.User{}).Find(req.Id, &user)
	resp.UserDetail = model.BuildUser(user)
	var vip model.Vip
	err = vip.FindActiveById(req.Id)
	resp.Vip = model.BuildVip(vip)
	return resp, err
}

func (*UserService) Create(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	user := &model.User{
		Account: req.Account,
	}
	model.DB.Create(user)
	resp = new(service.UserDetailResponse)
	resp.UserDetail = model.BuildUser(*user)
	return resp, nil
}

// func (*UserService) CreateByWechat(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
// 	return resp, nil
// }
// func (*UserService) Update(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
// 	return resp, nil
// }
