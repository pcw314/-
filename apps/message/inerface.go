package message

import (
	"gitee.com/xygfm/authorization/response"
	"github.com/gin-gonic/gin"
)

type Service interface {
	HandleMessage(ctx *gin.Context, msg Message)
	ListMessages(ctx *gin.Context, req *response.Paging, studentID int, enterpriseID int) ([]*Message, int64, error)
	ListConversation(ctx *gin.Context, role int, userID int) ([]*Conversation, error)
	ReadMessage(ctx *gin.Context, convID int, sender int) error
	// IncrementUnreadCount 增加指定用户在某会话中的未读数
	IncrementUnreadCount(ctx *gin.Context, convID, userID int) error
	// GetUnreadCount 获取指定用户在某会话中的未读数
	GetUnreadCount(ctx *gin.Context, convID, userID int) (int, error)
	// DecrementUnreadCount 减少指定用户在某会话中的未读数
	DecrementUnreadCount(ctx *gin.Context, convID, userID int) error
}

type ServiceImpl struct{}

func (s ServiceImpl) HandleMessage(ctx *gin.Context, msg Message) {

}
func (s ServiceImpl) ListMessages(ctx *gin.Context, req *response.Paging, studentID int, enterpriseID int) ([]*Message, int64, error) {
	return nil, 0, nil
}
func (s ServiceImpl) ListConversation(ctx *gin.Context, role int, userID int) ([]*Conversation, error) {
	return nil, nil
}
func (s ServiceImpl) ReadMessage(ctx *gin.Context, convID int, sender int) error {
	return nil
}
func (s ServiceImpl) IncrementUnreadCount(ctx *gin.Context, convID, userID int) error {
	return nil
}
func (s ServiceImpl) GetUnreadCount(ctx *gin.Context, convID, userID int) (int, error) {
	return 0, nil
}
func (s ServiceImpl) DecrementUnreadCount(ctx *gin.Context, convID, userID int) error {
	return nil
}
