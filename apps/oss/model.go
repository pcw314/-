package oss

import (
	utils "gitee.com/xygfm/authorization/util"
	"gorm.io/gorm"
)

type File struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Name      string `gorm:"column:name;comment:文件名" json:"name"`                          // 文件名
	Type      string `gorm:"column:type;comment:文件类型" json:"type"`                         // 文件类型
	Size      string `gorm:"column:size;default:0;comment:大小" json:"size"`                 // 大小
	CreatedAt int64  `gorm:"column:created_at;comment:上传时间" json:"created_at"`             // 上传时间
	Path      string `gorm:"column:path;comment:文件存储的路径" json:"path"`                      // 文件存储的路径
	OwnerType string `gorm:"column:owner_type" json:"owner_type"`
	OwnerID   int    `gorm:"column:owner_id;comment:上传的用户uid" json:"owner_id"` // 上传的用户uid
	State     int    `gorm:"column:state" json:"state"`                        // 状态：0关闭，1启用
}

func (*File) TableName() string {
	return "files"
}

// BeforeCreate 在创建之前执行的函数
func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	f.Path = utils.RemoveDomainName(f.Path)
	return
}

// BeforeUpdate 在更新之前执行的函数
func (f *File) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Path = utils.RemoveDomainName(f.Path)
	return
}

// AfterFind 在查询之后执行的函数
func (f *File) AfterFind(tx *gorm.DB) (err error) {
	api := utils.GetApiByType("file")
	f.Path = api.Url + f.Path
	return
}

type Api struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"` // 主键
	Type  string `gorm:"type" json:"type"`
	Url   string `gorm:"column:url" json:"url"`
	State int    `gorm:"column:state" json:"state"`
}

func (*Api) TableName() string {
	return "api"

}
