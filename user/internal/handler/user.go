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
	err = user.CheckPassword(req.Password)
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
	var platform model.UserPlatform
	l, _ := platform.SelectByUserId(req.Id)
	for _, up := range l {
		resp.Platform = append(resp.Platform, model.BuildUserPlatform(up))
	}
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

// func (*UserService) LoginByWechat(ctx context.Context, req *service.LoginByWechatRequest) (resp *service.UserDetailResponse, err error) {
// 	return resp, nil
// }

func (*UserService) Update(ctx context.Context, req *service.UserModel) (resp *service.Response, err error) {
	var user model.User
	resp = new(service.Response)
	model.DB.Model(&model.User{}).Find(req.Id, &user)
	err = user.SaveByService(req)
	return resp, err
}

func (*UserService) ModifyPassword(ctx context.Context, req *service.UserRequest) (resp *service.Response, err error) {
	var user model.User
	resp = new(service.Response)
	err = user.CheckUserIdExist(req.Id)
	return resp, err
}
