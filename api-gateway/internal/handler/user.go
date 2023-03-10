package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/e"
	"api-gateway/pkg/res"
	"api-gateway/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

// 用户登录
func (*UserService) Login(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.Login(context.Background(), &userReq)
	PanicIfUserError(err)
	token, _ := utils.GenerateToken(uint(userResp.UserDetail.Id))
	r := res.Response{
		Data: map[string]interface{}{
			"User":  userResp.UserDetail,
			"Token": token,
		},
		Code:  uint(userResp.Code),
		Error: e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func (*UserService) Create(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userResp, _ := userService.Create(context.Background(), &userReq)
	token, _ := utils.GenerateToken(uint(userResp.UserDetail.Id))
	r := res.Response{
		Data: map[string]interface{}{
			"User":  userResp.UserDetail,
			"Token": token,
		},
		Code:  uint(userResp.Code),
		Error: e.GetMsg(uint(userResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
func (*UserService) Update(ginCtx *gin.Context) {
	var userReq service.UserModel
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["user"].(service.UserServiceClient)
	userService.Update(context.Background(), &userReq)
	r := res.Response{}
	ginCtx.JSON(http.StatusOK, r)
}
