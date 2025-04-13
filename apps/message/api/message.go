package api

import (
	"encoding/json"
	"fmt"
	"gitee.com/xygfm/authorization/apps/message"
	"gitee.com/xygfm/authorization/response"
	"gitee.com/xygfm/authorization/response/result"
	utils "gitee.com/xygfm/authorization/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"log"
	"time"
)

func (h *handler) HandleWebSocket(c *gin.Context) {
	//convID, _ := c.GetQuery("conv_id")
	conn, err := message.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	userID := utils.GetUserID(c) // 从认证获取用户ID
	fmt.Println("userID:", userID)
	message.Mutex.Lock()
	// 检查并初始化会话对象
	//session, exists := message.Clients2[cast.ToUint(convID)]
	//if !exists {
	//	session = &message.Session{
	//		Connections: make(map[uint]*websocket.Conn),
	//	}
	//	message.Clients2[cast.ToUint(convID)] = session
	//}

	// 将连接绑定到用户
	//session.Connections[cast.ToUint(userID)] = conn
	message.Clients[cast.ToUint(userID)] = conn
	message.Mutex.Unlock()

	// 心跳检测
	//go h.handleHeartbeat(conn)

	//for {
	//	var msg message.Message
	//	if err := conn.ReadJSON(&msg); err != nil {
	//		log.Println("Error reading message:", err)
	//		// 清理连接
	//		message.Mutex.Lock()
	//		delete(message.Clients, cast.ToUint(userID))
	//		message.Mutex.Unlock()
	//		conn.Close()
	//		break
	//	}
	//	msg.SenderID = userID
	//	h.svc.HandleMessage(c, msg)
	//}
	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			// 清理连接
			message.Mutex.Lock()
			//delete(message.Clients2[cast.ToUint(convID)].Connections, cast.ToUint(userID))
			delete(message.Clients, cast.ToUint(userID))
			message.Mutex.Unlock()
			conn.Close()
			break
		}

		if messageType == websocket.TextMessage {
			var msg message.Message
			err := json.Unmarshal(messageBytes, &msg)
			if err != nil {
				log.Printf("Failed to parse JSON message: %v, raw data: %s", err, string(messageBytes))
				continue // 或者选择在这里关闭连接，取决于你的需求
			}
			msg.SenderID = userID
			h.svc.HandleMessage(c, msg)
		} else if messageType == websocket.CloseMessage {
			log.Println("Connection closed by client")
			message.Mutex.Lock()
			//delete(message.Clients2[cast.ToUint(convID)].Connections, cast.ToUint(userID))
			delete(message.Clients, cast.ToUint(userID))
			message.Mutex.Unlock()
			conn.Close()
			break
		}
		// 可能还需要处理其他类型的消息，如Ping/Pong
	}
}

// 心跳检测函数
func (h *handler) handleHeartbeat(conn *websocket.Conn) {
	fmt.Println("开启心跳检测")
	heartbeatInterval := 30 * time.Second
	pingTimeout := 10 * time.Second // 设定等待客户端回应的超时时间

	// 监听 Pong 消息，收到后刷新 ReadDeadline
	conn.SetPongHandler(func(appData string) error {
		log.Println("收到 Pong 响应")
		conn.SetReadDeadline(time.Now().Add(heartbeatInterval + pingTimeout))
		return nil
	})

	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 发送心跳包
			err := conn.WriteMessage(websocket.PingMessage, []byte("heartbeat"))
			if err != nil {
				log.Println("发送心跳失败，关闭连接:", err)
				conn.Close()
				return
			}

			// 设定超时，如果在 `pingTimeout` 时间内没有收到 Pong，连接会超时
			err = conn.SetReadDeadline(time.Now().Add(pingTimeout))
			if err != nil {
				log.Println("设置读取截止时间失败:", err)
				conn.Close()
				return
			}
		}
	}
}

func (h *handler) SendMessage(c *gin.Context) {
	var msg message.Message
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		return
	}
	msg.SenderID = utils.GetUserID(c)
	id := h.svc.HandleMessage(c, msg)
	response.Success(c, result.NewCorrect("发送成功", id))
}

func (h *handler) ListMessages(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	toUserID := cast.ToInt(ctx.Param("to_user_id"))
	convID := cast.ToInt(ctx.Param("conv_id"))
	userID := utils.GetUserID(ctx)
	err = h.svc.ReadMessage(ctx, convID, toUserID)
	if utils.GetUserRole(ctx) == 1 {
		if list, total, err := h.svc.ListMessages(ctx, &req, userID, toUserID, convID); err != nil {
			response.Error(ctx, result.DefaultError(err.Error()))
			return
		} else {
			response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
				Size:   req.Size,
				Page:   req.Page,
				List:   list,
				Total:  total,
				Search: req.Search,
			}))
		}
	} else if utils.GetUserRole(ctx) == 2 {
		if list, total, err := h.svc.ListMessages(ctx, &req, toUserID, userID, convID); err != nil {
			response.Error(ctx, result.DefaultError(err.Error()))
			return
		} else {
			response.Success(ctx, result.NewCorrect("获取成功", response.Paging{
				Size:   req.Size,
				Page:   req.Page,
				List:   list,
				Total:  total,
				Search: req.Search,
			}))
		}
	}
	return
}

func (h *handler) ListConversationWebSocket(ctx *gin.Context) {

	fmt.Println("34543534534")
	conn, err := message.Upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("22222222222")
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	userID := utils.GetUserID(ctx) // 从认证获取用户ID
	fmt.Println("userID:", userID)
	message.Mutex.Lock()
	message.ConvClients[cast.ToUint(userID)] = conn
	message.Mutex.Unlock()
	//go h.handleHeartbeat(conn)
	//conversation, err := h.svc.ListConversation(ctx, utils.GetUserRole(ctx), utils.GetUserID(ctx))
	//if err != nil {
	//	fmt.Println(err)
	//	response.Error(ctx, result.DefaultError(err.Error()))
	//	return
	//}
	//fmt.Println("23432423", conversation)
	//response.Success(ctx, result.NewCorrect("获取会话列表成功", conversation))
	return
}

func (h *handler) ListConversation(ctx *gin.Context) {
	date, _ := ctx.GetQuery("search")
	conversation, err := h.svc.ListConversation(ctx, utils.GetUserRole(ctx), utils.GetUserID(ctx), date)
	if err != nil {
		fmt.Println(err)
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取会话列表成功", conversation))
	return
}

func (h *handler) GetConversationID(ctx *gin.Context) {
	receiverID, _ := ctx.GetQuery("receiver_id")
	jobID, _ := ctx.GetQuery("job_id")
	id, err := h.svc.CreatedConversation(ctx, cast.ToInt(jobID), cast.ToInt(receiverID), utils.GetUserID(ctx), utils.GetUserRole(ctx))
	if err != nil {
		return
	}
	response.Success(ctx, result.NewCorrect("成功获取会话ID", id))
}

func (h *handler) ConversationTop(ctx *gin.Context) {

	convID := cast.ToInt(ctx.Param("conv_id"))
	err := h.svc.ConversationTop(ctx, convID, utils.GetUserRole(ctx))
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("置顶成功", ""))
	return
}

func (h *handler) DeleteConversationTop(ctx *gin.Context) {
	convID := cast.ToInt(ctx.Param("conv_id"))
	err := h.svc.DeleteConversationTop(ctx, convID, utils.GetUserRole(ctx))
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("取消成功", ""))
	return
}

func (h *handler) DeleteConversation(ctx *gin.Context) {
	convID := cast.ToInt(ctx.Param("conv_id"))
	err := h.svc.DeleteConversation(ctx, convID, utils.GetUserRole(ctx))
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("删除成功", ""))
	return
}
