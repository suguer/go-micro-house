package res

import (
	"api-gateway/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Data  interface{} `json:"Data"`
	Error string      `json:"Error"`
	Code  uint        `json:"Code"`
}

type PageResponse struct {
	Data       interface{} `json:"Data"`
	Error      string      `json:"Error"`
	Code       uint        `json:"Code"`
	Pagination interface{} `json:"Pagination"`
}

type Pagination struct {
	Total    int64 `json:"total" `
	Current  int   `json:"current" form:"current"`
	PageSize int   `json:"pageSize" form:"pageSize"`
}

func (m *Pagination) GetPageCurrent() int {
	if m.Current <= 0 {
		m.Current = 1
	}
	return m.Current
}

func (m *Pagination) GetPageSize() int {
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	return m.PageSize
}

// 返回200 自定义code data
func Ok(ctx *gin.Context, msgCode int, data interface{}) {
	ctx.JSON(http.StatusOK, ginH(msgCode, data))
}

// 无权限err
func Unauthorized(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusUnauthorized, ginH(msgCode, nil))
}

// 内部err
func InternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ginH(e.ERROR, nil))
}

// 禁止访问err
func ForbiddenError(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusForbidden, ginH(msgCode, nil))
}

// 自定义 err
func Error(ctx *gin.Context, httpCode, msgCode int) {
	ctx.JSON(httpCode, ginH(msgCode, nil))
}

func ginH(msgCode int, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(uint(msgCode)),
		"data": data,
	}
}
