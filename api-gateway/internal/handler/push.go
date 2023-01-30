package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"api-gateway/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PushService struct {
}

func (*PushService) Index(ginCtx *gin.Context) {
	var req service.SmsIndexRequest
	PanicIfUserError(ginCtx.Bind(&req))
	req.UserId = utils.GetUint32(ginCtx, "UserId")
	ser := ginCtx.Keys["sms"].(service.SmsServiceClient)
	resp, _ := ser.Index(context.Background(), &req)
	r := res.PageResponse{
		Data:       resp.Data,
		Pagination: resp.Pagination,
	}
	ginCtx.JSON(http.StatusOK, r)
}

func (*PushService) Create(ginCtx *gin.Context) {
	var req service.SmsRequest
	PanicIfUserError(ginCtx.Bind(&req))
	req.UserId = utils.GetUint32(ginCtx, "UserId")
	ser := ginCtx.Keys["sms"].(service.SmsServiceClient)
	resp, err := ser.Create(context.Background(), &req)
	if err != nil {
		ginCtx.JSON(http.StatusOK, res.Response{
			Error: err.Error(),
		})
		return
	}
	r := res.Response{
		Code: uint(resp.Code),
	}
	ginCtx.JSON(http.StatusOK, r)
}
