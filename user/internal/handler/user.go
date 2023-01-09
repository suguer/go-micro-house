package handler

import (
	"context"
	"user/internal/model"
	"user/internal/service"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var user model.User
	resp = new(service.UserDetailResponse)
	err = user.ShowUserInfo(req)
	if err != nil {
		return resp, err
	}
	resp.UserDetail = model.BuildUser(user)
	return resp, nil
}
