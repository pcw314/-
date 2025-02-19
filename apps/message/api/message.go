package api

import (
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
	conn, err := message.Upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	userID := utils.GetUserID(c) // 从认证获取用户ID
	fmt.Println("userID:", userID)
	message.Mutex.Lock()
	message.Clients[cast.ToUint(userID)] = conn
	message.Mutex.Unlock()

	// 心跳检测
	go h.handleHeartbeat(conn)

	for {
		var msg message.Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			// 清理连接
			message.Mutex.Lock()
			delete(message.Clients, cast.ToUint(userID))
			message.Mutex.Unlock()
			conn.Close()
			break
		}
		h.svc.HandleMessage(c, msg)
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

func (h *handler) ListMessages(ctx *gin.Context) {
	var req response.Paging
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	toUserID := cast.ToInt(ctx.Query("to_user_id"))
	convID := cast.ToInt(ctx.Query("conv_id"))
	userID := utils.GetUserID(ctx)
	err = h.svc.ReadMessage(ctx, convID, toUserID)
	if utils.GetUserRole(ctx) == 1 {
		if list, total, err := h.svc.ListMessages(ctx, &req, userID, toUserID); err != nil {
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
		if list, total, err := h.svc.ListMessages(ctx, &req, toUserID, userID); err != nil {
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

func (h *handler) ListConversation(ctx *gin.Context) {

	conn, err := message.Upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	userID := utils.GetUserID(ctx) // 从认证获取用户ID
	fmt.Println("userID:", userID)
	message.Mutex.Lock()
	message.ConvClients[cast.ToUint(userID)] = conn
	message.Mutex.Unlock()

	conversation, err := h.svc.ListConversation(ctx, utils.GetUserRole(ctx), utils.GetUserID(ctx))
	if err != nil {
		response.Error(ctx, result.DefaultError(err.Error()))
		return
	}
	response.Success(ctx, result.NewCorrect("获取会话列表成功", conversation))
	return
}
