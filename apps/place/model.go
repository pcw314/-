package place

type Area struct {
	ID       int     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Name     string  `gorm:"column:name" json:"name"`
	ParentID int     `gorm:"column:parent_id" json:"parent_id"` // 父菜单 ID
	Type     int     `gorm:"column:type" json:"type"`
	Pinyin   string  `gorm:"column:pinyin" json:"pinyin"`
	Zip      string  `gorm:"column:zip" json:"zip"`
	Sorting  int     `gorm:"column:sorting" json:"sorting"`
	GeoLong  float64 `gorm:"column:geo_long" json:"geo_long"`
	GeoLat   float64 `gorm:"column:geo_lat" json:"geo_lat"`
	State    int     `gorm:"column:state" json:"state"`
	IsSearch int     `gorm:"column:is_search" json:"is_search"`
	Children []*Area `gorm:"-" json:"children"`
}

func (Area) TableName() string {
	return "areas"
}

type School struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	ProvinceID int    `gorm:"column:province_id" json:"province_id"`
	Province   string `gorm:"-" json:"province"`
	CityID     int    `gorm:"column:city_id" json:"city_id"`
	City       string `gorm:"-" json:"city"`
	AreaID     int    `gorm:"column:area_id" json:"area_id"`
	Area       string `gorm:"-" json:"area"`
	State      int    `gorm:"column:state" json:"state"`
	VisitNum   int    `gorm:"column:visit_num" json:"visit_num"`
}

func (School) TableName() string {
	return "schools"
}
