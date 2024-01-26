package constant

const HeaderKeyAuthorization = "authorization"

const (
	NotificationTypeUserSearch = "user_search" // 玩家信息查询，玩家登录KOP商城时，需要用到玩家信息
	NotificationTypeServerList = "server_list" // 游戏服务器列表查询，玩家登录KOP商城时，可能需要用到游戏服务器ID
	NotificationTypeGift       = "gift"        // 赠品发放类型的消息通知，某些活动是不涉及资金流动的，例如：签到
	NotificationTypeOrder      = "order"       // 商城订单，当使用KOP商城时，会接收到此类型的消息通知
	NotificationTypePayment    = "payment"     // 交易
	NotificationTypeRefund     = "refund"      // 退款
	NotificationTypeDispute    = "dispute"     // 拒付
)

const (
	Success = "success"
	Failed  = "failed"
	Ignored = "ignored"
)

const (
	StatusPaid            = "PAID"
	StatusRefunded        = "REFUNDED"
	StatusDisputeAccepted = "DISPUTE_ACCEPTED"
)
