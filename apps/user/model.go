package user

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

type Student struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Sex    int8   `gorm:"column:sex" json:"sex"`
	Age    int    `gorm:"column:age" json:"age"`
	Major  string `gorm:"column:major" json:"major"` //专业
	Phone  string `gorm:"column:phone" json:"phone"` //手机号
	UserID int    `gorm:"column:user_id" json:"user_id"`
}

func (Student) TableName() string {
	return "students"
}

type Enterprise struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	LegaPerson string `gorm:"column:lega_person" json:"lega_person"` //法定代表人
	Phone      string `gorm:"column:phone" json:"phone"`
	State      int8   `gorm:"column:state" json:"state"`
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
	UserID     int    `gorm:"column:user_id" json:"user_id"`
}

func (Enterprise) TableName() string {
	return "enterprises"
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	ID          int    `json:"id"`
}
