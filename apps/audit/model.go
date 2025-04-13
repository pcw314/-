package audit

import (
	utils "gitee.com/xygfm/authorization/util"
	"gorm.io/gorm"
	"strings"
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
	if f.Image == "" {
		return nil
	}

	// 将图片 URL 按逗号分隔成数组
	imageURLs := strings.Split(f.Image, ",")

	// 创建一个新的切片，用于存储有效的图片 URL
	var validImageURLs []string

	// 遍历每个图片 URL，移除域名部分并过滤掉空字符串
	for _, imageURL := range imageURLs {
		trimmedURL := strings.TrimSpace(imageURL) // 去除前后空格
		if trimmedURL != "" {                     // 过滤掉空字符串
			validImageURLs = append(validImageURLs, utils.RemoveDomainName(trimmedURL))
		}
	}

	// 将处理后的图片 URL 重新拼接成字符串
	f.Image = strings.Join(validImageURLs, ",")
	return
}

// BeforeUpdate 在更新之前执行的函数
func (f *Audit) BeforeUpdate(tx *gorm.DB) (err error) {
	if f.Image == "" {
		return nil
	}

	// 将图片 URL 按逗号分隔成数组
	imageURLs := strings.Split(f.Image, ",")

	// 创建一个新的切片，用于存储有效的图片 URL
	var validImageURLs []string

	// 遍历每个图片 URL，移除域名部分并过滤掉空字符串
	for _, imageURL := range imageURLs {
		trimmedURL := strings.TrimSpace(imageURL) // 去除前后空格
		if trimmedURL != "" {                     // 过滤掉空字符串
			validImageURLs = append(validImageURLs, utils.RemoveDomainName(trimmedURL))
		}
	}

	// 将处理后的图片 URL 重新拼接成字符串
	f.Image = strings.Join(validImageURLs, ",")
	return
}

// AfterFind 在查询之后执行的函数
func (f *Audit) AfterFind(tx *gorm.DB) (err error) {
	api := utils.GetApiByType("file")
	if f.Image == "" {
		return nil
	}
	imageURLs := strings.Split(f.Image, ",")

	for i, imageURL := range imageURLs {
		imageURLs[i] = api.Url + strings.TrimSpace(imageURL)
	}
	f.Image = strings.Join(imageURLs, ",")

	return
}
