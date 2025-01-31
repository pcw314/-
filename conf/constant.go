package conf

var AdminMenu = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
var StudentMenu = []int{1, 2, 3, 4, 5}
var MerchantMenu = []int{1, 2, 3, 4, 5}

type Api struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Type  string `gorm:"type" json:"type"`
	Url   string `gorm:"column:url" json:"url"`
	State int    `gorm:"column:state" json:"state"`
}

func (*Api) TableName() string {
	return "api"
}
