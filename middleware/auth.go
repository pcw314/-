package service

import (
	"fmt"
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Security() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		fmt.Println("sfghskjhagfkjasfskbhdbxcxbcnabdsjad===========")
		claims, err := ParseToken(ctx.GetHeader("Authorization"))
		if err != nil {
			response.Fail(ctx, http.StatusUnauthorized, "Token校验失败")
			ctx.Abort()
			return
		}
		fmt.Println("claims:", claims)
		ctx.Next()
		return
	}
}

//func Security() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		// 打印所有请求头
//		fmt.Println("Headers:", ctx.Request.Header)
//
//		auth := ctx.GetHeader("Authorization")
//		fmt.Println("Auth header:", auth)
//
//		claims, err := ParseToken(auth)
//		if err != nil {
//			ctx.JSON(http.StatusUnauthorized, gin.H{
//				"error":   "unauthorized",
//				"message": err.Error(),
//			})
//			ctx.Abort()
//			return
//		}
//		fmt.Println("claims:", claims)
//		ctx.Next()
//	}
//}
