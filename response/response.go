package response

import (
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Base struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

//
//type result struct {
//	Code    int    `json:"code,omitempty"`
//	Message string `json:"message,omitempty"`
//	Data    any    `json:"data,omitempty"`
//}

type Page struct {
	PageSize *int        `json:"pageSize"`
	Current  *int        `json:"current"`
	List     interface{} `json:"list"`
	Total    *int64      `json:"total"`
}

type Paging struct {
	Size   int    `json:"size" form:"size"`
	Page   int    `json:"page" form:"page"`
	List   any    `json:"list" form:"list"`
	Total  int64  `json:"total" form:"total"`
	Search string `json:"search" form:"search"`
}

//func Success(ctx *gin.Context, code int, message string, data interface{}) {
//	ctx.JSON(http.StatusOK, result{
//		Code:    code,
//		Message: message,
//		Data:    data,
//	})
//}

func Fail(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusOK, result.Correct{
		Code:    code,
		Message: message,
	})
}
