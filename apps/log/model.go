package log

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
