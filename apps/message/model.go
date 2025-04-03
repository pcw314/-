package message

import (
	"gitee.com/xygfm/authorization/apps/user"
	utils "gitee.com/xygfm/authorization/util"
	"gorm.io/gorm"
)

type Conversation struct {
	ID             int              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EnterpriseID   int              `gorm:"column:enterprise_id" json:"enterprise_id"`
	EnterpriseInfo *user.Enterprise `gorm:"-" json:"enterprise_info"`
	StudentID      int              `gorm:"column:student_id" json:"student_id"`
	StudentInfo    *user.Student    `gorm:"-" json:"student_info"`
	JobID          int              `gorm:"column:job_id" json:"job_id"`
	State          string           `gorm:"column:state" json:"state"`
	UpdateAt       int64            `gorm:"column:update_at" json:"update_at"`
	CreatedAt      int64            `gorm:"column:created_at" json:"created_at"`
	StudentTop     int              `gorm:"column:student_top" json:"student_top"`
	EnterpriseTop  int              `gorm:"column:enterprise_top" json:"enterprise_top"`
	IsDeleted      int              `gorm:"column:is_deleted" json:"is_deleted"`
	IsTop          int              `gorm:"-" json:"is_top"`
	UnReadNum      int              `gorm:"-" json:"unread_num"`
	Msg            string           `gorm:"-" json:"msg"`
	ContactName    string           `gorm:"-" json:"contact_name"`
	Post           string           `gorm:"-" json:"post"`
}

func (Conversation) TableName() string {
	return "conversations"
}

type Message struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ConvID     int    `gorm:"column:conv_id;primaryKey;autoIncrement:true" json:"conv_id"`
	SenderID   int    `gorm:"column:sender_id;primaryKey;autoIncrement:true" json:"sender_id"`
	ReceiverID int    `gorm:"-" json:"receiver_id"`
	Content    string `gorm:"column:content" json:"content"`
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`
	IsRead     int    `gorm:"column:is_read" json:"is_read"`
}

func (Message) TableName() string {
	return "messages"
}

type Enterprise struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Phone  string `gorm:"column:phone" json:"phone"`
	Avatar string `gorm:"colum:avatar" json:"avatar"`
	UserID int    `gorm:"column:user_id" json:"user_id"`
}
type Student struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	UserID int    `gorm:"column:user_id" json:"user_id"`
	Avatar string `gorm:"colum:avatar" json:"avatar"`
}
type ConversationList struct {
	ID             int         `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EnterpriseID   int         `gorm:"column:enterprise_id" json:"enterprise_id"`
	EnterpriseInfo *Enterprise `gorm:"-" json:"enterprise_info"`
	StudentID      int         `gorm:"column:student_id" json:"student_id"`
	StudentInfo    *Student    `gorm:"-" json:"student_info"`
	JobID          int         `gorm:"column:job_id" json:"job_id"`
	UpdateAt       int64       `gorm:"column:update_at" json:"update_at"`
	StudentTop     int         `gorm:"column:student_top" json:"student_top"`
	EnterpriseTop  int         `gorm:"column:enterprise_top" json:"enterprise_top"`
	IsTop          int         `gorm:"-" json:"is_top"`
	UnReadNum      int         `gorm:"-" json:"unread_num"`
	Msg            string      `gorm:"-" json:"msg"`
	ContactName    string      `gorm:"-" json:"contact_name"`
	Post           string      `gorm:"-" json:"post"`
}

func (f *Enterprise) AfterFind(tx *gorm.DB) (err error) {
	api := utils.GetApiByType("file")
	if f.Avatar != "" {
		f.Avatar = api.Url + f.Avatar
	}
	return
}
func (f *Student) AfterFind(tx *gorm.DB) (err error) {
	api := utils.GetApiByType("file")
	if f.Avatar != "" {
		f.Avatar = api.Url + f.Avatar
	}
	return
}
