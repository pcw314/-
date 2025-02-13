package admin

type User struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Username     string `gorm:"colum:username" json:"username"`
	Password     string `gorm:"colum:password" json:"password"`
	Role         int    `gorm:"colum:role" json:"role"`
	HeadPortrait string `gorm:"colum:head_portrait" json:"head_portrait"`
}

func (User) TableName() string {
	return "users"
}

type LoginRequest struct {
	Username string `gorm:"colum:username" json:"username"`
	Password string `gorm:"colum:password" json:"password"`
	Role     int    `gorm:"colum:role" json:"role"`
}

func (LoginRequest) TableName() string {
	return "users"
}

type Student struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Name  string `gorm:"colum:name" json:"name"`
	Sex   int8   `gorm:"colum:sex" json:"sex"` //性别  0未设置，1男  2女
	Age   int    `gorm:"colum:age" json:"age"`
	Major string `gorm:"colum:major" json:"major"`
	Phone string `gorm:"colum:phone" json:"phone"`
}

func (Student) TableName() string {
	return "students"
}

// MenuRequest 菜单
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

type Areas struct {
	ID       int      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string   `gorm:"colum:name" json:"name"`
	ParentID int      `gorm:"column:parent_id" json:"parent_id"`
	Type     int8     `gorm:"colum:type" json:"type"`
	Pinyin   string   `gorm:"colum:pinyin" json:"pinyin"`
	Zip      string   `gorm:"colum:zip" json:"zip"`
	Sorting  int      `gorm:"colum:sorting" json:"sorting"`
	GeoLong  float32  `gorm:"colum:geo_long" json:"geo_long"`
	GeoLat   float32  `gorm:"colum:geo_lat" json:"geo_lat"`
	State    int8     `gorm:"colum:state" json:"state"`
	IsSearch int8     `gorm:"colum:is_search" json:"is_search"`
	Children []*Areas `gorm:"-" json:"children"`
}

func (Areas) TableName() string {
	return "areas"
}

type Menu struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	ParentID int     `gorm:"column:parent_id" json:"parent_id"`
	Type     string  `gorm:"column:type" json:"type"` // 类型：menu 或 button
	Method   string  `gorm:"column:method" json:"method"`
	Title    string  `gorm:"column:title" json:"title"` // 标题
	Path     string  `gorm:"column:path" json:"path"`   // 路由路径
	Icon     string  `gorm:"column:icon" json:"icon"`   // 图标
	Sort     int     `gorm:"column:sort" json:"sort"`
	IsHide   int8    `gorm:"column:is_hide" json:"is_hide"`
	Children []*Menu `gorm:"-" json:"children"`
}

func (Menu) TableName() string {
	return "menus"
}
