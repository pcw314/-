package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/message"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"time"
)

func (i *impl) HandleMessage(ctx *gin.Context, msg message.Message) {
	// 存储到MySQL
	if msg.ConvID == 0 {
		if utils.GetUserRole(ctx) == 1 {
			conversation := message.Conversation{
				EnterpriseID: msg.ReceiverID,
				StudentID:    utils.GetUserID(ctx),
				State:        "open",
				CreatedAt:    time.Now().UnixMilli(),
			}
			err := i.mdb.Model(&message.Conversation{}).
				Create(&conversation).Error
			if err != nil {
				fmt.Println(err)
				return
			}
			msg.ConvID = conversation.ID
		} else if utils.GetUserRole(ctx) == 2 {
			conversation := message.Conversation{
				EnterpriseID: utils.GetUserID(ctx),
				StudentID:    msg.ReceiverID,
				State:        "open",
				CreatedAt:    time.Now().UnixMilli(),
			}
			err := i.mdb.Model(&message.Conversation{}).
				Create(&conversation).Error
			if err != nil {
				fmt.Println(err)
				return
			}
			msg.ConvID = conversation.ID
		}
	}
	var role int
	i.mdb.Model(&user.User{}).Select("role").Where("id = ?", msg.ReceiverID).First(&role)
	fmt.Println("53405495", role)
	fmt.Println("msg", msg)
	i.mdb.Model(&message.Message{}).Create(&message.Message{
		ConvID:    msg.ConvID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		CreatedAt: time.Now().UnixMilli(),
	})

	// 实时推送
	message.Mutex.Lock()
	receiverConn, ok := message.Clients[uint(msg.ReceiverID)]
	message.Mutex.Unlock()

	if ok {
		err := receiverConn.WriteJSON(gin.H{
			"type": "new_message",
			"data": msg,
		})
		if err != nil {
			fmt.Println("向客户端发送消息失败:", err)
			receiverConn.Close()                      // 关闭失败的连接
			delete(message.Clients, uint(msg.IsRead)) // 移除客户端
			return
		}
		fmt.Println("=============")
		fmt.Println(message.Clients)
	} else {
		fmt.Println("452345423423423")
		err := i.IncrementUnreadCount(ctx, msg.ConvID, msg.ReceiverID)
		if err != nil {
			fmt.Println("增加未读数失败", err)
		}
		fmt.Println(i.GetUnreadCount(ctx, msg.ConvID, msg.ReceiverID))
		convList, err := i.ListConversation(ctx, role, msg.ReceiverID)
		if err != nil {
			return
		}
		message.Mutex.Lock()
		coveConn, ok := message.ConvClients[uint(msg.ReceiverID)]
		message.Mutex.Unlock()
		if ok {
			err := coveConn.WriteJSON(gin.H{
				"type": "new_convList",
				"data": convList,
			})
			if err != nil {
				fmt.Println("向客户端发送消息失败:", err)
				coveConn.Close()                              // 关闭失败的连接
				delete(message.ConvClients, uint(msg.IsRead)) // 移除客户端
				return
			}
		}

	}
}

//	func (i *impl) HandleUnreadMessage(ctx *gin.Context) {
//		var convs []message.Conversation
//		userID := utils.GetUserID(ctx)
//		i.mdb.Preload()
//	}
func (i *impl) ListMessages(ctx *gin.Context, req *response.Paging, studentID int, enterpriseID int) ([]*message.Message, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var ConversationID int
	err := i.mdb.Model(&message.Conversation{}).
		Select("id").
		Where("enterprise_id = ?", enterpriseID).
		Where("student_id = ?", studentID).
		First(&ConversationID).Error
	if err != nil {
		return nil, 0, err
	}
	var pos []*message.Message
	db := i.mdb.Model(&message.Message{}).Where("conv_id = ?", ConversationID)
	if req.Search != "" {
		db = db.Where("content LIKE ?", "%"+req.Search+"%")
	}
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&pos).Error
	return pos, 0, nil
}

func (i *impl) ReadMessage(ctx *gin.Context, convID int, sender int) error {
	err := i.mdb.Model(&message.Message{}).
		Where("conv_id = ?", convID).
		Where("sender_id = ?", sender).
		Update("is_read", 1).
		Error
	if err != nil {
		return err
	}
	err = i.DecrementUnreadCount(ctx, convID, utils.GetUserID(ctx))
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) ListConversation(ctx *gin.Context, role int, userID int) ([]*message.Conversation, error) {
	var pos []*message.Conversation
	db := i.mdb.Model(&message.Conversation{})
	if role == 1 {
		db.Where("student_id = ?", userID)
	} else if role == 2 {
		db.Where("enterprise_id = ?", userID)
	}
	err := db.Find(&pos).Error
	if err != nil {
		return nil, err
	}
	var userIDs []int
	if role == 1 {
		for _, item := range pos {
			userIDs = append(userIDs, item.EnterpriseID)
		}
		var enterprise []*user.Enterprise
		enterpriseMAP := make(map[int]*user.Enterprise)
		err = i.mdb.Model(&user.Enterprise{}).
			Where("user_id in ?", userIDs).
			Find(&enterprise).Error
		if err != nil {
			return nil, err
		}
		for _, item := range enterprise {
			enterpriseMAP[item.UserID] = item
		}
		for j, item := range pos {
			pos[j].UnReadNum, _ = i.GetUnreadCount(ctx, item.ID, item.StudentID)
			pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
		}
	} else if role == 2 {
		for _, item := range pos {
			userIDs = append(userIDs, item.StudentID)
		}
		var student []*user.Student
		studentMAP := make(map[int]*user.Student)
		err = i.mdb.Model(&user.Student{}).
			Where("user_id in ?", userIDs).
			Find(&student).Error
		if err != nil {
			return nil, err
		}
		for _, item := range student {
			studentMAP[item.UserID] = item
		}
		for j, item := range pos {
			pos[j].UnReadNum, _ = i.GetUnreadCount(ctx, item.ID, item.EnterpriseID)
			pos[j].StudentInfo = studentMAP[item.EnterpriseID]
		}
	}
	return pos, nil
}

func (i *impl) IncrementUnreadCount(ctx *gin.Context, convID, userID int) error {
	key := fmt.Sprintf("conv:%d:user:%d:unread", convID, userID)
	return i.rdb.Incr(ctx, key).Err()
}
func (i *impl) GetUnreadCount(ctx *gin.Context, convID, userID int) (int, error) {
	key := fmt.Sprintf("conv:%d:user:%d:unread", convID, userID)
	return i.rdb.Get(ctx, key).Int()
}
func (i *impl) DecrementUnreadCount(ctx *gin.Context, convID, userID int) error {
	key := fmt.Sprintf("conv:%d:user:%d:unread", convID, userID)
	return i.rdb.Decr(ctx, key).Err()
}
