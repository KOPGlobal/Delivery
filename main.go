package main

import (
	"context"
	"delivery/constant"
	"delivery/games"
	"delivery/games/params"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeliveryHandler(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": constant.Failed,
			"error":   err.Error(),
		})
		return
	}

	// 初步解析
	var notification params.Notification
	err = json.Unmarshal(body, &notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": constant.Failed,
			"error":   err.Error(),
		})
		return
	}

	// 验签
	ctx := context.Background()
	sign := c.GetHeader(constant.HeaderKeyAuthorization)
	project := games.NewProject(notification.ProjectId)
	if project == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": constant.Failed,
			"error":   "project error",
		})
		return
	}

	// 验签失败
	if !project.CheckSign(ctx, body, sign) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": constant.Failed,
			"error":   "invalid signature",
		})
		return
	}

	// 处理消息通知
	var res params.CommonResult
	switch notification.NotificationType {
	case constant.NotificationTypeOrder:
		// order
		res = project.HandleOrder(ctx, notification)
	case constant.NotificationTypeGift:
		// gift
		res = project.HandleGift(ctx, notification)
	case constant.NotificationTypeUserSearch:
		// user_search
		res = project.SearchUser(ctx, notification)
	case constant.NotificationTypeServerList:
		// server_list
		res = project.ListServers(ctx, notification)
	case constant.NotificationTypePayment:
		// payment
		res = project.HandlePayment(ctx, notification)
	case constant.NotificationTypeRefund:
		// refund
		res = project.HandleRefund(ctx, notification)
	case constant.NotificationTypeDispute:
		//dispute
		res = project.HandleDispute(ctx, notification)
	default:
		res.HttpStatusCode = http.StatusOK
		res.Message = constant.Ignored
	}

	c.JSON(res.HttpStatusCode, gin.H{
		"message": res.Message,
		"data":    res.Data,
		"error":   res.Error,
	})
}

func main() {
	r := gin.Default()

	// 定义一个GET请求的处理函数
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// 接收KOP发来的消息通知
	r.POST("/notify", DeliveryHandler)

	// 启动服务
	r.Run(":8080")
}
