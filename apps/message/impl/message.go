package impl

import (
	"fmt"
	"gitee.com/xygfm/authorization/apps/message"
	"gitee.com/xygfm/authorization/apps/position"
	"gitee.com/xygfm/authorization/apps/user"
	"gitee.com/xygfm/authorization/response"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"time"
)

func (i *impl) HandleMessage(ctx *gin.Context, msg message.Message) int {
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
				return 0
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
				return 0
			}
			msg.ConvID = conversation.ID
		}
	}
	var role int
	i.mdb.Model(&user.User{}).Select("role").Where("id = ?", msg.ReceiverID).First(&role)
	fmt.Println("53405495", role)
	fmt.Println("msg", msg)
	mess := &message.Message{
		ConvID:    msg.ConvID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		CreatedAt: time.Now().UnixMilli(),
	}
	err := i.mdb.Model(&message.Message{}).Create(&mess).Error
	//更新会话时间
	err = i.SetConversationUpdateAt(ctx, msg.ConvID)
	if err != nil {
		return 0
	}

	// 实时推送
	message.Mutex.Lock()
	//receiverConn, ok := message.Clients2[uint(msg.ConvID)].Connections[uint(msg.ReceiverID)]
	receiverConn, ok := message.Clients[uint(msg.ReceiverID)]
	message.Mutex.Unlock()

	if ok {
		err = receiverConn.WriteJSON(gin.H{
			"type": "new_message",
			"data": mess,
		})
		if err != nil {
			fmt.Println("向客户端发送消息失败:", err)
			receiverConn.Close() // 关闭失败的连接
			//delete(message.Clients2[uint(msg.ConvID)].Connections, uint(msg.IsRead))
			delete(message.Clients, uint(msg.IsRead)) // 移除客户端
			return 0
		}
		i.mdb.Model(&message.Message{}).Where("id = ?", mess.ID).Updates(map[string]interface{}{
			"is_read": 1,
		})

	} else {
		err = i.IncrementUnreadCount(ctx, msg.ConvID, msg.ReceiverID)
		if err != nil {
			fmt.Println("增加未读数失败", err)
		}
		fmt.Println(i.GetUnreadCount(ctx, msg.ConvID, msg.ReceiverID))
		convList, err := i.ListConversation(ctx, role, msg.ReceiverID, "")
		if err != nil {
			return 0
		}
		message.Mutex.Lock()
		coveConn, ok := message.ConvClients[uint(msg.ReceiverID)]
		message.Mutex.Unlock()
		if ok {
			err = coveConn.WriteJSON(gin.H{
				"type": "new_convList",
				"data": convList,
			})
			if err != nil {
				fmt.Println("向客户端发送消息失败:", err)
				coveConn.Close()                              // 关闭失败的连接
				delete(message.ConvClients, uint(msg.IsRead)) // 移除客户端
				return 0
			}
		}

	}
	return mess.ID
}

//	func (i *impl) HandleUnreadMessage(ctx *gin.Context) {
//		var convs []message.Conversation
//		userID := utils.GetUserID(ctx)
//		i.mdb.Preload()
//	}
func (i *impl) ListMessages(ctx *gin.Context, req *response.Paging, studentID int, enterpriseID int, convID int) ([]*message.Message, int64, error) {
	limit := req.Size
	offset := (req.Page - 1) * limit
	var ConversationID int
	if convID != 0 {
		ConversationID = convID
	} else {
		err := i.mdb.Model(&message.Conversation{}).
			Select("id").
			Where("enterprise_id = ?", enterpriseID).
			Where("student_id = ?", studentID).
			First(&ConversationID).Error
		if err != nil {
			return nil, 0, err
		}
	}
	var pos []*message.Message
	db := i.mdb.Model(&message.Message{}).Where("conv_id = ?", ConversationID)
	if req.Search != "" {
		db = db.Where("content LIKE ?", "%"+req.Search+"%")
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Order("created_at desc").Find(&pos).Error
	reverseMessages(pos)
	return pos, total, nil
}
func reverseMessages(messages []*message.Message) {
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
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

//	func (i *impl) ListConversation(ctx *gin.Context, role int, userID int, date string) ([]*message.Conversation, error) {
//		var pos []*message.Conversation
//		db := i.mdb.Model(&message.Conversation{})
//		if date != "" {
//			db = db.Where("update_at >= ?", time.Now().AddDate(0, 0, -30).UnixMilli())
//		}
//		if role == 1 {
//			db = db.Where("student_id = ?", userID)
//			// 添加基于student_top的排序，并且在相同student_top值的情况下，按update_at降序排序
//			db = db.Order("student_top desc, update_at desc")
//		} else if role == 2 {
//			db = db.Where("enterprise_id = ?", userID)
//			// 添加基于enterprise_top的排序，并且在相同enterprise_top值的情况下，按update_at降序排序
//			db = db.Order("enterprise_top desc, update_at desc")
//		} else {
//			// 如果没有匹配的角色，则仅按update_at降序排序
//			db = db.Order("update_at desc")
//		}
//
//		err := db.Find(&pos).Error
//		if err != nil {
//			return nil, err
//		}
//		var userIDs []int
//		var convIDs []int
//		var jobIDs []int
//		if role == 1 {
//			for _, item := range pos {
//				userIDs = append(userIDs, item.EnterpriseID)
//				convIDs = append(convIDs, item.ID)
//				jobIDs = append(jobIDs, item.JobID)
//			}
//			var studentInfo *user.Student
//			i.mdb.Model(&user.Student{}).Where("user_id = ?", userID).First(&studentInfo)
//			var enterprise []*user.Enterprise
//			enterpriseMAP := make(map[int]*user.Enterprise)
//			err = i.mdb.Model(&user.Enterprise{}).
//				Where("user_id in ?", userIDs).
//				Find(&enterprise).Error
//			if err != nil {
//				return nil, err
//			}
//			for _, item := range enterprise {
//				enterpriseMAP[item.UserID] = item
//			}
//			var msg []*message.Message
//			msgMap := make(map[int]string)
//
//			subQuery := i.mdb.Table("messages").Select("conv_id, MAX(created_at) as max_created_at").
//				Where("conv_id IN (?)", convIDs).
//				Group("conv_id")
//			err = i.mdb.Table("messages").Joins("JOIN (?) AS latest ON messages.conv_id = latest.conv_id AND messages.created_at = latest.max_created_at", subQuery).
//				Where("messages.conv_id IN (?)", convIDs).
//				Find(&msg).Error
//			for _, item := range msg {
//				msgMap[item.ConvID] = item.Content
//			}
//
//			var job []*position.Job
//			contactNameMap := make(map[int]string)
//			postMap := make(map[int]string)
//			err = i.mdb.Model(&position.Job{}).
//				Select("id,contact_name, post").
//				Where("id in (?)", jobIDs).
//				Find(&job).Error
//			if err != nil {
//				return nil, err
//			}
//			for _, item := range job {
//				contactNameMap[item.ID] = item.ContactName
//				postMap[item.ID] = item.Post
//			}
//			for j, item := range pos {
//				pos[j].UnReadNum, _ = i.GetUnreadCount(ctx, item.ID, item.StudentID)
//				pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
//				pos[j].Msg = msgMap[item.ID]
//				pos[j].ContactName = contactNameMap[item.JobID]
//				pos[j].Post = postMap[item.JobID]
//				pos[j].IsTop = pos[j].StudentTop
//				pos[j].StudentInfo = studentInfo
//			}
//		} else if role == 2 {
//			for _, item := range pos {
//				userIDs = append(userIDs, item.StudentID)
//				convIDs = append(convIDs, item.ID)
//				jobIDs = append(jobIDs, item.JobID)
//			}
//			var enterpriseInfo *user.Enterprise
//			i.mdb.Model(&user.Enterprise{}).Where("user_id = ?", userID).First(&enterpriseInfo)
//			var student []*user.Student
//			studentMAP := make(map[int]*user.Student)
//			err = i.mdb.Model(&user.Student{}).
//				Where("user_id in ?", userIDs).
//				Find(&student).Error
//			if err != nil {
//				return nil, err
//			}
//			for _, item := range student {
//				studentMAP[item.UserID] = item
//			}
//			var msg []*message.Message
//			msgMap := make(map[int]string)
//			subQuery := i.mdb.Table("messages").Select("conv_id, MAX(created_at) as max_created_at").
//				Where("conv_id IN (?)", convIDs).
//				Group("conv_id")
//			err = i.mdb.Table("messages").Joins("JOIN (?) AS latest ON messages.conv_id = latest.conv_id AND messages.created_at = latest.max_created_at", subQuery).
//				Where("messages.conv_id IN (?)", convIDs).
//				Find(&msg).Error
//			if err != nil {
//				return nil, err
//			}
//			for _, item := range msg {
//				msgMap[item.ConvID] = item.Content
//			}
//
//			var job []*position.Job
//			contactNameMap := make(map[int]string)
//			postMap := make(map[int]string)
//			err = i.mdb.Model(&position.Job{}).
//				Select("id,contact_name,post").
//				Where("id in (?)", jobIDs).
//				Find(&job).Error
//			if err != nil {
//				return nil, err
//			}
//			for _, item := range job {
//				contactNameMap[item.ID] = item.ContactName
//				postMap[item.ID] = item.Post
//			}
//
//			for j, item := range pos {
//				pos[j].UnReadNum, _ = i.GetUnreadCount(ctx, item.ID, item.EnterpriseID)
//				pos[j].StudentInfo = studentMAP[item.StudentID]
//				pos[j].Msg = msgMap[item.ID]
//				pos[j].ContactName = contactNameMap[item.JobID]
//				pos[j].Post = postMap[item.JobID]
//				pos[j].IsTop = pos[j].EnterpriseTop
//				pos[j].EnterpriseInfo = enterpriseInfo
//			}
//		}
//		return pos, nil
//	}
func (i *impl) ListConversation(ctx *gin.Context, role int, userID int, date string) ([]*message.ConversationList, error) {
	var pos []*message.ConversationList
	db := i.mdb.Model(&message.Conversation{})
	if date != "" {
		db = db.Where("update_at >= ?", time.Now().AddDate(0, 0, -30).UnixMilli())
	}
	if role == 1 {
		db = db.Where("student_id = ?", userID).Where("is_deleted != ?", 1)

		// 添加基于student_top的排序，并且在相同student_top值的情况下，按update_at降序排序
		db = db.Order("student_top desc, update_at desc")
	} else if role == 2 {
		db = db.Where("enterprise_id = ?", userID).Where("is_deleted != ?", 2)

		// 添加基于enterprise_top的排序，并且在相同enterprise_top值的情况下，按update_at降序排序
		db = db.Order("enterprise_top desc, update_at desc")
	} else {
		// 如果没有匹配的角色，则仅按update_at降序排序
		db = db.Order("update_at desc")
	}

	err := db.Find(&pos).Error
	if err != nil {
		return nil, err
	}
	var userIDs []int
	var convIDs []int
	var jobIDs []int
	if role == 1 {
		for _, item := range pos {
			userIDs = append(userIDs, item.EnterpriseID)
			convIDs = append(convIDs, item.ID)
			jobIDs = append(jobIDs, item.JobID)
		}
		var studentInfo *message.Student
		i.mdb.Model(&user.Student{}).Where("user_id = ?", userID).First(&studentInfo)
		var enterprise []*message.Enterprise
		enterpriseMAP := make(map[int]*message.Enterprise)
		err = i.mdb.Model(&user.Enterprise{}).
			Where("user_id in ?", userIDs).
			Find(&enterprise).Error
		if err != nil {
			return nil, err
		}
		for _, item := range enterprise {
			enterpriseMAP[item.UserID] = item
		}
		var msg []*message.Message
		msgMap := make(map[int]string)

		subQuery := i.mdb.Table("messages").Select("conv_id, MAX(created_at) as max_created_at").
			Where("conv_id IN (?)", convIDs).
			Group("conv_id")
		err = i.mdb.Table("messages").Joins("JOIN (?) AS latest ON messages.conv_id = latest.conv_id AND messages.created_at = latest.max_created_at", subQuery).
			Where("messages.conv_id IN (?)", convIDs).
			Find(&msg).Error
		for _, item := range msg {
			msgMap[item.ConvID] = item.Content
		}

		var job []*position.Job
		contactNameMap := make(map[int]string)
		postMap := make(map[int]string)
		err = i.mdb.Model(&position.Job{}).
			Select("id,contact_name, post").
			Where("id in (?)", jobIDs).
			Find(&job).Error
		if err != nil {
			return nil, err
		}
		for _, item := range job {
			contactNameMap[item.ID] = item.ContactName
			postMap[item.ID] = item.Post
		}
		for j, item := range pos {
			pos[j].UnReadNum, _ = i.GetUnreadCount(ctx, item.ID, item.StudentID)
			pos[j].EnterpriseInfo = enterpriseMAP[item.EnterpriseID]
			pos[j].Msg = msgMap[item.ID]
			pos[j].ContactName = contactNameMap[item.JobID]
			pos[j].Post = postMap[item.JobID]
			pos[j].IsTop = pos[j].StudentTop
			pos[j].StudentInfo = studentInfo
		}
	} else if role == 2 {
		for _, item := range pos {
			userIDs = append(userIDs, item.StudentID)
			convIDs = append(convIDs, item.ID)
			jobIDs = append(jobIDs, item.JobID)
		}
		var enterpriseInfo *message.Enterprise
		i.mdb.Model(&user.Enterprise{}).Where("user_id = ?", userID).First(&enterpriseInfo)
		var student []*message.Student
		studentMAP := make(map[int]*message.Student)
		err = i.mdb.Model(&user.Student{}).
			Where("user_id in ?", userIDs).
			Find(&student).Error
		if err != nil {
			return nil, err
		}
		for _, item := range student {
			studentMAP[item.UserID] = item
		}
		var msg []*message.Message
		msgMap := make(map[int]string)
		subQuery := i.mdb.Table("messages").Select("conv_id, MAX(created_at) as max_created_at").
			Where("conv_id IN (?)", convIDs).
			Group("conv_id")
		err = i.mdb.Table("messages").Joins("JOIN (?) AS latest ON messages.conv_id = latest.conv_id AND messages.created_at = latest.max_created_at", subQuery).
			Where("messages.conv_id IN (?)", convIDs).
			Find(&msg).Error
		if err != nil {
			return nil, err
		}
		for _, item := range msg {
			msgMap[item.ConvID] = item.Content
		}

		var job []*position.Job
		postMap := make(map[int]string)
		err = i.mdb.Model(&position.Job{}).
			Select("id,contact_name,post").
			Where("id in (?)", jobIDs).
			Find(&job).Error
		if err != nil {
			return nil, err
		}
		for _, item := range job {
			postMap[item.ID] = item.Post
		}

		for j, item := range pos {
			// 获取未读消息计数
			unreadCount, err := i.GetUnreadCount(ctx, item.ID, item.EnterpriseID)
			if err != nil {
				// 处理错误，例如记录日志或设置默认值
				unreadCount = 0
			}
			pos[j].UnReadNum = unreadCount
			// 检查 studentMAP 中是否存在对应的 StudentID
			if studentInfo, exists := studentMAP[item.StudentID]; exists {
				pos[j].StudentInfo = studentInfo
				pos[j].ContactName = studentInfo.Name
			} else {
				// 设置默认值或处理缺失的情况
				pos[j].StudentInfo = nil // 假设 Student 是你的学生信息结构体
				pos[j].ContactName = "未知"
			}
			pos[j].Msg = msgMap[item.ID]
			pos[j].Post = postMap[item.JobID]
			pos[j].IsTop = pos[j].EnterpriseTop
			pos[j].EnterpriseInfo = enterpriseInfo
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
	return i.rdb.Set(ctx, key, 0, 0).Err()
}

func (i *impl) CreatedConversation(ctx *gin.Context, jobID, receiverID, senderID, role int) (int, error) {
	var id int
	var err error
	if role == 1 {
		err = i.mdb.Model(&message.Conversation{}).
			Select("id").
			Where("student_id = ?", senderID).
			Where("enterprise_id = ?", receiverID).
			Where("job_id = ?", jobID).
			First(&id).Error

	} else if role == 2 {
		err = i.mdb.Model(&message.Conversation{}).
			Select("id").
			Where("student_id = ?", receiverID).
			Where("enterprise_id = ?", senderID).
			Where("job_id = ?", jobID).
			First(&id).Error
	}
	if err != nil {
		if role == 1 {
			conversation := message.Conversation{
				EnterpriseID: receiverID,
				StudentID:    senderID,
				JobID:        jobID,
				State:        "open",
				UpdateAt:     time.Now().UnixMilli(),
				CreatedAt:    time.Now().UnixMilli(),
				IsDeleted:    0,
			}
			err = i.mdb.Model(&message.Conversation{}).
				Create(&conversation).Error
			if err != nil {
				fmt.Println(err)
				return 0, err
			}
			id = conversation.ID
		} else if role == 2 {
			conversation := message.Conversation{
				EnterpriseID: senderID,
				StudentID:    receiverID,
				JobID:        jobID,
				State:        "open",
				UpdateAt:     time.Now().UnixMilli(),
				CreatedAt:    time.Now().UnixMilli(),
				IsDeleted:    0,
			}
			err = i.mdb.Model(&message.Conversation{}).
				Create(&conversation).Error
			if err != nil {
				fmt.Println(err)
				return 0, err
			}
			id = conversation.ID
		}

	}

	return id, nil
}

func (i *impl) SetConversationUpdateAt(ctx *gin.Context, id int) error {
	err := i.mdb.Model(&message.Conversation{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"update_at": time.Now().UnixMilli()}).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *impl) ConversationTop(ctx *gin.Context, convID, role int) error {
	if role == 1 {
		err := i.mdb.Model(&message.Conversation{}).
			Where("id = ?", convID).
			Updates(map[string]interface{}{"student_top": 1}).Error
		if err != nil {
			return err
		}
	} else if role == 2 {
		err := i.mdb.Model(&message.Conversation{}).
			Where("id = ?", convID).
			Updates(map[string]interface{}{"enterprise_top": 1}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) DeleteConversationTop(ctx *gin.Context, convID, role int) error {
	if role == 1 {
		err := i.mdb.Model(&message.Conversation{}).
			Where("id = ?", convID).
			Updates(map[string]interface{}{"student_top": 0}).Error
		if err != nil {
			return err
		}
	} else if role == 2 {
		err := i.mdb.Model(&message.Conversation{}).
			Where("id = ?", convID).
			Updates(map[string]interface{}{"enterprise_top": 0}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) DeleteConversation(ctx *gin.Context, convID, role int) error {
	fmt.Println("convID", convID)
	fmt.Println("role", role)
	var conv *message.Conversation
	i.mdb.Model(&message.Conversation{}).Where("id = ?", convID).First(&conv)
	if role == 1 {
		if conv.IsDeleted == 2 {
			i.mdb.Model(&message.Conversation{}).Where("id = ?", convID).Delete(&message.Conversation{})
		} else {
			i.mdb.Model(&message.Conversation{}).Where("id = ?", convID).Updates(map[string]interface{}{"is_deleted": 1})
		}
	} else if role == 2 {
		if conv.IsDeleted == 1 {
			i.mdb.Model(&message.Conversation{}).Where("id = ?", convID).Delete(&message.Conversation{})
		} else {
			i.mdb.Model(&message.Conversation{}).Where("id = ?", convID).Updates(map[string]interface{}{"is_deleted": 2})
		}
	}
	return nil
}
