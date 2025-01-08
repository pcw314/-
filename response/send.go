package response

import (
	"gitee.com/xygfm/authorization/response/result"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, cor *result.Correct) {
	ctx.JSON(cor.Code, cor)
}

func Error(ctx *gin.Context, err *result.Error) {
	ctx.JSON(err.Code, err)
}
