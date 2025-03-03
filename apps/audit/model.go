package audit

import (
	utils "gitee.com/xygfm/authorization/util"
	"gorm.io/gorm"
)

type Audit struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Informer     int    `gorm:"column:informer" json:"informer"`                   //举报人
	InformerName string `gorm:"-" json:"informer_name"`
	Reason       string `gorm:"column:reason" json:"reason"`
	Image        string `gorm:"column:image" json:"image"`
	UserID       int    `gorm:"column:user_id" json:"user_id"` //被举报人
	Name         string `gorm:"-" json:"name"`
	Phone        string `gorm:"-" json:"phone"`
	Avatar       string `gorm:"-" json:"avatar"`
	JobID        int    `gorm:"column:job_id" json:"job_id"`
	JobInfo      *Job   `gorm:"-" json:"job_info"`
	Reply        string `gorm:"column:reply" json:"reply"`
	State        int    `gorm:"column:state" json:"state"` //0：待审核  1：已审核
	Auditor      int    `gorm:"auditor" json:"auditor"`
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at" json:"updated_at"`
}

type AuditRequest struct {
	Reason string `gorm:"column:reason" json:"reason"`
	Image  string `gorm:"column:image" json:"image"`
	UserID int    `gorm:"column:user_id" json:"user_id"` //被举报人
	JobID  int    `gorm:"column:job_id" json:"job_id"`
}

type UserInfo struct {
	Name   string `gorm:"-" json:"name"`
	Phone  string `gorm:"-" json:"phone"`
	Avatar string `gorm:"-" json:"avatar"`
}

func (Audit) TableName() string {
	return "audits"
}

type Job struct {
	ID            int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EnterpriseID  int    `gorm:"column:enterprise_id" json:"enterprise_id"`
	Name          string `gorm:"column:name" json:"name"`
	Type          string `gorm:"column:type" json:"type"`
	WorkingTime   string `gorm:"column:working_time" json:"working_time"`
	Place         string `gorm:"column:place" json:"place"`
	Salary        string `gorm:"column:salary" json:"salary"`
	Post          string `gorm:"column:post" json:"post"`
	SchoolID      int    `gorm:"column:school_id" json:"school_id"`
	Description   string `gorm:"column:description" json:"description"`
	Require       string `gorm:"column:require" json:"require"`
	ContactName   string `gorm:"column:contact_name" json:"contact_name"`
	ContactNumber string `gorm:"column:contact_number" json:"contact_number"`
	State         int    `gorm:"column:state" json:"state"`
	CreatedAt     int64  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at" json:"updated_at"`
}

// BeforeCreate 在创建之前执行的函数
func (f *Audit) BeforeCreate(tx *gorm.DB) (err error) {
	f.Image = utils.RemoveDomainName(f.Image)
	return
}

// BeforeUpdate 在更新之前执行的函数
func (f *Audit) BeforeUpdate(tx *gorm.DB) (err error) {
	f.Image = utils.RemoveDomainName(f.Image)
	return
}

// AfterFind 在查询之后执行的函数
func (f *Audit) AfterFind(tx *gorm.DB) (err error) {
	api := utils.GetApiByType("file")
	if f.Image != "" {
		f.Image = api.Url + f.Image
	}
	return
}
