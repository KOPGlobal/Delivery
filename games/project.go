package games

import (
	"context"
	"delivery/games/params"
)

type Project interface {
	CheckSign(context.Context, []byte, string) bool                         // 验签
	SearchUser(context.Context, params.Notification) params.CommonResult    // 处理查询玩家信息类型的消息通知，玩家在KOP商城中绑定玩家ID时，会查询玩家信息
	ListServers(context.Context, params.Notification) params.CommonResult   // 处理查询游戏服务器列表的消息通知，玩家在KOP商城中绑定玩家服务器ID时，可能会用到游戏服务器ID
	HandlePayment(context.Context, params.Notification) params.CommonResult // 处理交易类型的消息通知
	HandleRefund(context.Context, params.Notification) params.CommonResult  // 处理退款类型的消息通知
	HandleDispute(context.Context, params.Notification) params.CommonResult // 处理拒付类型的消息通知
	HandleOrder(context.Context, params.Notification) params.CommonResult   // 处理商城订单类型的消息通知
	HandleGift(context.Context, params.Notification) params.CommonResult    // 处理赠品发放类型的消息通知，某些活动例如签到，玩家参与活动后会发送纯粹礼品发放的消息通知
}

func NewProject(projectId int64) Project {
	if projectId == SampleProjectId {
		return newSample()
	}

	return nil
}
