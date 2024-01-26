package games

import (
	"context"
	"crypto/sha1"
	"delivery/constant"
	"delivery/games/params"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SampleProjectId     = 123456
	SampleWebhookSecret = "sample_webhook_secret"
)

// Sample 示例游戏
type Sample struct{}

func newSample() *Sample {
	return &Sample{}
}

func (s *Sample) getProjectId(ctx context.Context) int64 {
	return SampleProjectId
}

func (s *Sample) getWebhookSecret(ctx context.Context) string {
	return SampleWebhookSecret
}

func (s *Sample) sign(ctx context.Context, payload []byte) string {
	data := string(payload)
	data += s.getWebhookSecret(ctx)
	h := sha1.New()
	h.Write([]byte(data))
	return fmt.Sprintf("Signature %s", hex.EncodeToString(h.Sum(nil)))
}

func (s *Sample) CheckSign(ctx context.Context, payload []byte, sign string) bool {
	return s.sign(ctx, payload) == sign
}

// ship 发放道具
func (s *Sample) ship(ctx context.Context, order params.Order) params.CommonResult {
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}
	///////////////////////////////////////////////////
	// TODO 给玩家发放道具
	fmt.Println("ship order: ", fmt.Sprintf("%+v", order))
	///////////////////////////////////////////////////
	return res
}

// deduct 扣除道具
func (s *Sample) deduct(ctx context.Context, order params.Order) params.CommonResult {
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}
	///////////////////////////////////////////////////
	// TODO 扣除玩家的道具
	fmt.Println("deduct order: ", fmt.Sprintf("%+v", order))
	///////////////////////////////////////////////////
	return res
}

// sendGift 发放赠品
func (s *Sample) sendGift(ctx context.Context, gift params.Gift) params.CommonResult {
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}
	///////////////////////////////////////////////////
	// TODO 给玩家发放赠品
	fmt.Println("sendGift gift: ", fmt.Sprintf("%+v", gift))
	///////////////////////////////////////////////////
	return res
}

// SearchUser 查询玩家信息，玩家在KOP商城购买道具时，需要绑定玩家id，才能下单
func (s *Sample) SearchUser(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
		玩家信息查询类型的消息通知的数据结构
		{
		  "notification_type": "user_search",
		  "settings": {
		    "merchant_id": 1006805,
		    "project_id": 1006842
		  },
		  "user": {
		    "id": "3234",             // 玩家id
		    "server": "S1",           // 玩家服务器id
		    "merchant_id": 1006805,
		    "project_id": 1006842
		  }
		}
	*/
	///////////////////////////////////////////////////
	// TODO 从数据库里面查询玩家信息
	fmt.Println("SearchUser notification: ", fmt.Sprintf("%+v", notification))
	var user = params.User{
		Id:      "xxxxxxxxxxxxx",
		Name:    "xxxxx",
		Level:   10,
		Server:  "xxx",
		Country: "XX",
	}
	///////////////////////////////////////////////////
	data := params.UserSearchResponseData{
		User: user,
	}
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           data,
		Message:        constant.Success,
		Error:          "",
	}
	return res
}

// ListServers 查询游戏服务器列表
// 玩家在KOP商城购物时，需要同时绑定玩家id和玩家服务器id时，才需要实现此方法
// 如果不需要绑定服务器id，则不需要实现此方法
func (s *Sample) ListServers(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
		服务器列表查询类型消息通知的数据结构
		{
		  "notification_type": "server_list",
		  "settings": {
		    "merchant_id": 1006805,
		    "project_id": 1006791
		  },
		  "user": {
		    "merchant_id": 1006805,
		    "project_id": 1006791
		  }
		}
	*/
	///////////////////////////////////////////////////
	// TODO 查询服务器列表
	fmt.Println("ListServers notification: ", fmt.Sprintf("%+v", notification))
	data := params.ServerListResponseData{
		Servers: nil,
	}
	for i := 1; i < 10; i++ {
		data.Servers = append(data.Servers, params.Server{
			Id:   fmt.Sprintf("%d", i),
			Name: fmt.Sprintf("server-%d", i),
		})
	}
	///////////////////////////////////////////////////
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           data,
		Message:        constant.Success,
		Error:          "",
	}
	return res
}

// HandleOrder 处理KOP商城订单
func (s *Sample) HandleOrder(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
			KOP商城-订单类型消息通知数据结构
		{
		  "event_id": "2bc3a80f944e9777c31235344c7868d3",
		  "notification_type": "order",
		  "settings": {
		    "merchant_id": 1000001,
		    "project_id": 1000002
		  },
		  "data": {
		    "amount": 1,
		    "country": "BR",
		    "create_time": 1684152939,
		    "currency": "BRL",
		    "language": "cn",
		    "merchant_id": 5,
		    "metadata": {
		      "gifts": [
		        {
		          "gift_goods_id": "12345",
		          "logo": "https://kopglobal.com/1678755548457_\u91d1\u5e0101.png",
		          "num": 1
		        }
		      ],
		      "items": [
		        {
		          "amount": 1,
		          "code": "884002",
		          "currency": "BRL",
		          "id": 271,
		          "names": {
		            "en": "Monta d'Oro dellaterra desolata*5000",
		            "lang_id": "Monta d'Oro dellaterra desolata*5000"
		          },
		          "num": 1,
		          "order_id": 956900,
		          "price": 1,
		          "purchase_limits": 0
		        }
		      ],
		      "order_id": "100100001965",
		      "sale_biz_hit": [
		        {
		          "activity_rule": "{\"goods_list\":[{\"goods_id\":271,\"num\":1}],\"gift_list\":[{\"gift_goods_id\":\"12345\",\"num\":1,\"gift_goods_logo\":\"https://kopglobal.com/site/prod/store/1678755548457_\u91d1\u5e0101.png\"}]}",
		          "activity_type": 1,
		          "activity_type_txt": "买X赠Y",
		          "biz_id": 49,
		          "can_gift": true,
		          "gift_list": [
		            {
		              "gift_goods_id": "12345",
		              "logo": "https://kopglobal.com/1678755548457_\u91d1\u5e0101.png",
		              "num": 1
		            }
		          ],
		          "name": "1",
		          "project_id": 1000002,
		          "start_time": 1678697932,
		          "stop_time": 4822948800
		        }
		      ]
		    },
		    "mode": "live",
		    "payer": {
		        "id": "123456",  // 玩家游戏中的uid，必传
		        "server": "S1",  // 玩家所在服务器id，非必传
		    },
		    "payer_id": "123456",
		    "payment_type": 1,
		    "project_id": 1000002,
		    "rfc_1766_language": "zh-CN",
		    "status": "PAID",
		    "transaction_id": "23051510001930"
		  }
		}

	*/
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}

	///////////////////////////////////////////////////
	// TODO 从 notification 中解析出KOP商城订单信息
	fmt.Println("HandleOrder notification: ", fmt.Sprintf("%+v", notification))

	var order params.Order
	data, _ := json.Marshal(&notification.Data)
	_ = json.Unmarshal(data, &order)

	// 处理 order
	switch order.Status {
	case constant.StatusPaid:
		// 处理付款成功的订单
		// 给玩家发放道具
		res = s.ship(ctx, order)
		//
		// TODO 其他操作
		//

	case constant.StatusRefunded:
		// 处理退款的订单
		// 扣除玩家的道具
		res = s.deduct(ctx, order)
		//
		// TODO 其他操作
		//

	case constant.StatusDisputeAccepted:
		// 处理拒付的订单
		// 扣除玩家的道具
		res = s.deduct(ctx, order)
		//
		// TODO 其他操作
		//

	default:
		// TODO 其他状态的订单，可以忽略
		res.Message = constant.Ignored
	}

	///////////////////////////////////////////////////
	return res
}

// HandleGift 处理赠品发放相关的消息通知
func (s *Sample) HandleGift(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
			发放赠品类型的消息通知数据结构
		{
		  "event_id": "",
		  "project_id": 1000002,
		  "notification_type": "gift",
		  "data": {
		    "merchant_id": 1000001,
		    "project_id": 1000002,
		    "player_id": "tsb-1777777",
		    "server_id": "S2",
		    "mode": "",
		    "biz_info": {
		      "id": 452,
		      "name": "测试签到",
		      "type": 4,
		      "start_time": 1706237987,
		      "Stop_time": 1706324387,
		      "biz_type": "activity"
		    },
		    "gift_info": {
		      "gift_list": [
		        {
		          "gift_goods_id": "1212121",
		          "num": 1,
		          "gift_goods_logo": "https://kopglobal.com/8d8027490711352f8673cdd890a25f0d.jpeg"
		        }
		      ]
		    },
		    "create_time": 1706239355,
		    "language": "",
		    "rfc_1766_language": "",
		    "metadata": {
		      "sign_in_rule_type": 2,
		      "is_circulate": false,
		      "sign_in_total_days": 1,
		      "free_appending_days": 0,
		      "gift_list": [
		        {
		          "gift_goods_id": "1212121",
		          "num": 1,
		          "gift_goods_logo": "https://kopglobal.com/8d8027490711352f8673cdd890a25f0d.jpeg",
		          "day_no": 1,
		          "is_sign_in": false,
		          "is_available_sign_in": false,
		          "is_appending": false
		        }
		      ],
		      "consecutive_gift_list": [

		      ]
		    }
		  }
		}
	*/
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}

	///////////////////////////////////////////////////
	// TODO 从 notification 解析出 gift 并处理 发放赠品
	fmt.Println("HandleGift notification: ", fmt.Sprintf("%+v", notification))

	var gift params.Gift
	data, _ := json.Marshal(&notification.Data)
	_ = json.Unmarshal(data, &gift)

	// 处理 gift
	res = s.sendGift(ctx, gift)

	///////////////////////////////////////////////////
	return res
}

// HandlePayment 处理支付相关的消息通知
func (s *Sample) HandlePayment(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
		交易-支付类型消息通知的数据结构
		{
		  "event_id": "d5e2487cf4e442369a9381e65958afd2",
		  "notification_type": "payment",
		  "settings": {
		    "merchant_id": 1000001,
		    "project_id": 1000002
		  },
		  "data": {
		    "amount": 8,
		    "country": "HK",
		    "create_time": 1691657752,
		    "currency": "HKD",
		    "merchant_id": 1000001,
		    "merchant_transaction_id": "600100004072",
		    "metadata": {
		      "a": "1",
		      "b": "2"
		    },
		    "mode": "sandbox",
		    "payer": {
		      "id": "player-FD1SETRGBVC"
		    },
		    "project_id": 1000002,
		    "status": "PAID",
		    "transaction_id": "23081060004113"
		  }
		}
	*/
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}

	///////////////////////////////////////////////////
	// TODO 从 notification 解析出并处理 支付信息
	fmt.Println("HandlePayment notification: ", fmt.Sprintf("%+v", notification))

	var payment params.Payment
	data, _ := json.Marshal(&notification.Data)
	_ = json.Unmarshal(data, &payment)

	// 处理payment
	switch payment.Status {
	case constant.StatusPaid:
		// TODO 处理此状态的支付即可
	default:
		// TODO 其他状态的支付，可以忽略
		res.Message = constant.Ignored
	}

	///////////////////////////////////////////////////
	return res
}

// HandleRefund 处理退款
func (s *Sample) HandleRefund(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
		交易-退款类型消息数据结构
		{
		    "event_id": "123abcdef123abcdef",           // 事件id
		    "notification_type": "refund",              // 消息类型
		    "data": {
		        "create_time": 177123456,               // 创建时间
		        "amount": 100,                          // 金额
		        "currency": "RUB",                      // 币种
		        "reason": "",                           // 退款原因
		        "refund_id": "4232323232323",           // 退款id
		        "transaction_id": "64232323232323",     // 原始交易id
		        "status": "REFUNDED",                   // 退款状态
		        "update_time": 177123456,               // 更新时间
		        "metadata": {
		            "a": "1",
		            "b": "2"
		        },                                      // 商户自定义数据
		        "merchant_transaction_id": "123456",    // 商户端交易id
		        "merchant_id": 1000001,                 // 商户id
		        "project_id": 1000002,                  // 应用id
		        "payer": {
		            "id": "51232323323232"              // 商户端游戏内玩家id
		        },
		        "mode": "sandbox"                       // live正式环境，sandbox沙盒环境
		    }
		}
	*/
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}

	///////////////////////////////////////////////////
	// TODO 从 notification 解析出refund 并处理退款
	fmt.Println("HandleRefund notification: ", fmt.Sprintf("%+v", notification))

	var refund params.Refund
	data, _ := json.Marshal(&notification.Data)
	_ = json.Unmarshal(data, &refund)

	// 处理 refund
	switch refund.Status {
	case constant.StatusRefunded:
		// TODO 处理此状态的退款即可
	default:
		// TODO 其他状态的退款，可以忽略
		res.Message = constant.Ignored
	}

	///////////////////////////////////////////////////
	return res
}

// HandleDispute 处理拒付
func (s *Sample) HandleDispute(ctx context.Context, notification params.Notification) params.CommonResult {
	/*
		交易-拒付类型消息通知数据结构
		{
		    "event_id": "123abcdef123abcdef",           // 事件id
		    "notification_type": "dispute",             // 消息类型
		    "data": {
		        "create_time": 177123456,               // 创建时间
		        "amount": 100,                          // 金额
		        "currency": "RUB",                      // 币种
		        "reason": "",                           // 退款原因
		        "dispute_id": "4232323232323",          // 拒付id
		        "transaction_id": "64232323232323",     // 原始交易id
		        "status": "DISPUTE_ACCEPTED",           // 拒付状态
		        "update_time": 177123456,               // 更新时间
		        "metadata": {
		            "a": "1",
		            "b": "2"
		        },                                      // 商户自定义数据
		        "merchant_transaction_id": "123456",    // 商户端交易id
		        "merchant_id": 1000001,                      // 商户id
		        "project_id": 1000002,                       // 应用id
		        "payer": {
		            "id": "51232323323232"              // 商户端游戏内玩家id
		        },
		        "mode": "sandbox"                       // live正式环境，sandbox沙盒环境
		    }
		}
	*/
	res := params.CommonResult{
		HttpStatusCode: http.StatusOK,
		Data:           nil,
		Message:        constant.Success,
		Error:          "",
	}

	///////////////////////////////////////////////////
	// TODO 从 notification 解析出 dispute 并处理 拒付
	fmt.Println("HandleDispute notification: ", fmt.Sprintf("%+v", notification))

	var dispute params.Dispute
	data, _ := json.Marshal(&notification.Data)
	_ = json.Unmarshal(data, &dispute)

	// 处理 dispute
	switch dispute.Status {
	case constant.StatusDisputeAccepted:
		// TODO 处理此状态的拒付即可
	default:
		// TODO 其他状态的拒付，可以忽略
		res.Message = constant.Ignored
	}

	///////////////////////////////////////////////////
	return res
}
