package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HouseGroupService struct {
}

func (*HouseGroupService) Index(ginCtx *gin.Context) {
	var req service.HouseGroupIndexRequest
	PanicIfUserError(ginCtx.Bind(&req))
	ser := ginCtx.Keys["house_group"].(service.HouseGroupServiceClient)
	resp, _ := ser.Index(context.Background(), &req)
	r := res.Response{
		Code: uint(resp.Code),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func (*HouseGroupService) Create(ginCtx *gin.Context) {
	var req service.HouseGroupRequest
	PanicIfUserError(ginCtx.Bind(&req))
	ser := ginCtx.Keys["house_group"].(service.HouseGroupServiceClient)
	resp, _ := ser.Create(context.Background(), &req)
	r := res.Response{
		Data: resp.HouseGroupModel,
		Code: uint(resp.Code),
	}
	ginCtx.JSON(http.StatusOK, r)
}
