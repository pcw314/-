package conf

type Api struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Type  string `gorm:"type" json:"type"`
	Url   string `gorm:"column:url" json:"url"`
	State int    `gorm:"column:state" json:"state"`
}

func (*Api) TableName() string {
	return "api"
}
