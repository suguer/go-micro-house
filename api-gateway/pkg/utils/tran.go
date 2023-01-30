package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUint32(ginCtx *gin.Context, str string) uint32 {
	g, _ := ginCtx.Get(str)
	v, _ := strconv.Atoi(g.(string))
	return uint32(v)
}
