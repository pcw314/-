package message

import "gitee.com/xygfm/authorization/apps/user"

type Conversation struct {
	ID             int              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EnterpriseID   int              `gorm:"column:enterprise_id" json:"enterprise_id"`
	EnterpriseInfo *user.Enterprise `gorm:"-" json:"enterprise_info"`
	StudentID      int              `gorm:"column:student_id" json:"student_id"`
	StudentInfo    *user.Student    `gorm:"-" json:"student_info"`
	State          string           `gorm:"column:state" json:"state"`
	CreatedAt      int64            `gorm:"column:created_at" json:"created_at"`
	UnReadNum      int              `gorm:"-" json:"unread_num"`
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
