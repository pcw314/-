package utils

import (
	service "gitee.com/xygfm/authorization/middleware"
	"github.com/gin-gonic/gin"
)

func GetContextClaims(ctx *gin.Context, key string) *service.UserClaims {
	value, exists := ctx.Get(key)
	if !exists {
		return nil
	}
	// 尝试将 interface{} 转换为 UserClaims
	if claims, ok := value.(*service.UserClaims); ok {
		return claims
	}
	return nil
}

func GetUserID(ctx *gin.Context) int {
	claims := GetContextClaims(ctx, "claims")
	if claims == nil {
		return 0
	}
	return claims.UserId
}

func GetUserRole(ctx *gin.Context) int {
	claims := GetContextClaims(ctx, "claims")
	if claims == nil {
		return 0
	}
	return claims.Role
}
