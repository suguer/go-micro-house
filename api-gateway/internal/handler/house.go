package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HouseService struct {
}

func (*HouseService) Index(ginCtx *gin.Context) {
	var req service.HouseIndexRequest
	PanicIfUserError(ginCtx.Bind(&req))
	ser := ginCtx.Keys["house"].(service.HouseServiceClient)
	resp, _ := ser.Index(context.Background(), &req)
	r := res.Response{
		Data: resp.Data,
	}
	ginCtx.JSON(http.StatusOK, r)
}

func (*HouseService) GroupIndex(ginCtx *gin.Context) {
	var req service.HouseGroupRequest
	PanicIfUserError(ginCtx.Bind(&req))
	ser := ginCtx.Keys["house_group"].(service.HouseGroupServiceClient)
	resp, _ := ser.Index(context.Background(), &req)
	r := res.Response{
		Code: int(resp.Code),
	}
	ginCtx.JSON(http.StatusOK, r)
}
