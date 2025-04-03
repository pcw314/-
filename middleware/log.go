package service

import (
	"fmt"
	"gitee.com/xygfm/authorization/global"
	"github.com/gin-gonic/gin"
	"time"
)

type Log struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Method    string  `gorm:"column:method" json:"method"`                       // 请求方法
	Path      string  `gorm:"column:path" json:"path"`                           // 请求路径
	Query     string  `gorm:"column:query" json:"query"`                         // 查询参数
	Status    int     `gorm:"column:status" json:"status"`                       // HTTP状态码
	IP        string  `gorm:"column:ip" json:"ip"`                               // 客户端IP地址
	Latency   float64 `gorm:"column:latency" json:"latency"`                     // 延迟时间（毫秒）
	UserAgent string  `gorm:"column:user_agent" json:"user_agent"`               // 用户代理
	Timestamp int64   `gorm:"column:timestamp" json:"timestamp"`                 // 时间戳
}

func (Log) TableName() string {
	return "logs"
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start).Seconds() * 1000 // 转换为毫秒

		logEntry := Log{
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			Query:     c.Request.URL.RawQuery,
			Status:    c.Writer.Status(),
			IP:        c.ClientIP(),
			Latency:   latency,
			UserAgent: c.GetHeader("User-Agent"),
			Timestamp: time.Now().UnixMilli(),
		}

		// 使用GORM进行插入操作
		if err := global.Mdb.Model(&Log{}).Create(&logEntry).Error; err != nil {
			fmt.Println("Failed to insert log into database:", err)
		}
	}
}
