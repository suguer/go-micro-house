package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HouseService struct {
}

func (*HouseService) Index(ginCtx *gin.Context) {
	var req service.HouseIndexRequest
	PanicIfUserError(ginCtx.Bind(&req))
	ser := ginCtx.Keys["house"].(service.HouseServiceClient)
	resp, err := ser.Index(context.Background(), &req)
	fmt.Printf("err: %v\n", err)
	r := res.PageResponse{
		Data:       resp.Data,
		Pagination: resp.Pagination,
	}
	ginCtx.JSON(http.StatusOK, r)
}
