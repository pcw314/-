package template

type Template struct {
	ID  int `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	PID int `gorm:"column:pid" json:"pid"`                             // 父菜单 ID
}

func (Template) TableName() string {
	return "Templates"
}
