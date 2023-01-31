package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 200 //操作成功
	FAILED  int = 404 //操作失败
)

//普通成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"smg":  "成功",
		"data": v,
	})
}

//普通的操作失败返回
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}
