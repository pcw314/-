package position

import (
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
)

type Job struct {
	ID             int              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EnterpriseID   int              `gorm:"column:enterprise_id" json:"enterprise_id"`
	EnterpriseInfo *user.Enterprise `gorm:"-" json:"enterprise_info"`
	UserInfo       *user.User       `gorm:"-" json:"user_info"`
	Type           int              `gorm:"column:type" json:"type"`
	WorkingTime    string           `gorm:"column:working_time" json:"working_time"`
	Place          string           `gorm:"column:place" json:"place"`
	Salary         string           `gorm:"column:salary" json:"salary"`
	Post           string           `gorm:"column:post" json:"post"`
	SchoolID       int              `gorm:"column:school_id" json:"school_id"`
	Description    string           `gorm:"column:description" json:"description"`
	Require        string           `gorm:"column:require" json:"require"`
	ContactName    string           `gorm:"column:contact_name" json:"contact_name"`
	ContactNumber  string           `gorm:"column:contact_number" json:"contact_number"`
	State          int              `gorm:"column:state" json:"state"`
	Reply          string           `gorm:"column:reply" json:"reply"`
	CreatedAt      int64            `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      int64            `gorm:"column:updated_at" json:"updated_at"`
	IsCollect      int              `gorm:"-" json:"is_collect"`
}

func (Job) TableName() string {
	return "jobs"
}

type Collect struct {
	ID        int   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int   `gorm:"column:user_id" json:"user_id"`
	JobID     int   `gorm:"column:job_id" json:"job_id"`
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"`
}

func (Collect) TableName() string {
	return "collects"
}

type AuditJobRequest struct {
	response.Paging
	State int `json:"state" form:"state"`
}

type UpdateAuditJob struct {
	ID    int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	State int    `gorm:"-" json:"state"`
	Reply string `gorm:"-" json:"reply"`
}
