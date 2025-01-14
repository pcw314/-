package user

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (LoginRequest) TableName() string {
	return "store_menus_authority"
}

type MenuRequest struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	PID       int            `gorm:"column:pid" json:"pid"`                             // 父菜单 ID
	Type      string         `gorm:"column:type" json:"type"`                           // 类型：menu 或 button
	Title     string         `gorm:"column:title" json:"title"`                         // 标题
	Name      string         `gorm:"column:name" json:"name"`                           // 名称
	Path      string         `gorm:"column:path" json:"path"`                           // 路由路径
	Icon      string         `gorm:"column:icon" json:"icon"`                           // 图标
	MenuType  string         `gorm:"column:menu_type" json:"menu_type"`                 // 菜单类型（如 tab）
	URL       string         `gorm:"column:url" json:"url"`                             // URL
	Component string         `gorm:"column:component" json:"component"`                 // 组件路径
	KeepAlive string         `gorm:"column:keepalive" json:"keepalive"`                 // 是否缓存
	Extend    string         `gorm:"column:extend" json:"extend"`                       // 扩展字段
	Children  []*MenuRequest `gorm:"-" json:"children"`
}

func (MenuRequest) TableName() string {
	return "permissions"
}
