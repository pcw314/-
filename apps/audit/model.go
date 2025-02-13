package audit

import (
	"gitee.com/xygfm/authorization/apps/user"
)

type Audit struct {
	ID             int              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 唯一标识符
	Reason         string           `gorm:"column:reason" json:"reason"`
	Image          string           `gorm:"column:image" json:"image"`
	EnterpriseID   int              `gorm:"column:enterprise_id" json:"enterprise_id"`
	EnterpriseInfo *user.Enterprise `gorm:"-" json:"enterprise_info"`
	JobID          int              `gorm:"column:job_id" json:"job_id"`
	JobInfo        *Job             `gorm:"-" json:"job_info"`
	Reply          string           `gorm:"column:reply" json:"reply"`
	State          int              `gorm:"column:state" json:"state"` //0：待审核  1：已审核
	Auditor        int              `gorm:"auditor" json:"auditor"`
	CreatedAt      int64            `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      int64            `gorm:"column:updated_at" json:"updated_at"`
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
