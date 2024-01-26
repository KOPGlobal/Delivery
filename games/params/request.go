package params

type CommonResult struct {
	HttpStatusCode int         // http状态码
	Data           interface{} // 具体的数据
	Message        string      // 描述
	Error          string      // 错误信息
}

type Card struct {
	Bin         string `json:"bin"`
	Brand       string `json:"brand"`
	CardType    string `json:"card_type"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	Last4       string `json:"last4"`
}

type Payer struct {
	Id     string `json:"id"`               // 玩家id
	Name   string `json:"name,omitempty"`   // 玩家名字
	Server string `json:"server,omitempty"` // 玩家服务器id
	Email  string `json:"email,omitempty"`  // 玩家email
	Card   *Card  `json:"card,omitempty"`   // 信用卡信息
}

type Payment struct {
	Amount                float64                `json:"amount"`                  // 支付金额
	Country               string                 `json:"country"`                 // 国家
	CreateTime            int64                  `json:"create_time"`             // 创建时间
	Currency              string                 `json:"currency"`                // 币种
	MerchantId            int64                  `json:"merchant_id"`             // 商户id
	MerchantTransactionId string                 `json:"merchant_transaction_id"` // 商户端订单id
	Metadata              map[string]interface{} `json:"metadata"`                // 源数据，创建交易时传给KOP，交易成功，再由KOP回传给商户
	Mode                  string                 `json:"mode"`                    // 交易模式 沙盒 sandbox， 生产 live
	PayId                 string                 `json:"pay_id"`                  // 交易id
	Payer                 Payer                  `json:"payer"`                   // 付款人
	ProjectId             int64                  `json:"project_id"`              // 应用id
	Status                string                 `json:"status"`                  // 交易状态
	TransactionId         string                 `json:"transaction_id"`          // 交易id
	UpdateTime            int64                  `json:"update_time"`             // 更新时间
}

type Refund struct {
	Amount                float64                `json:"amount"`                  // 退款金额
	CreateTime            int64                  `json:"create_time"`             // 创建时间
	Currency              string                 `json:"currency"`                // 币种
	MerchantId            int64                  `json:"merchant_id"`             // 商户id
	MerchantTransactionId string                 `json:"merchant_transaction_id"` // 商户端退款id
	Metadata              map[string]interface{} `json:"metadata"`                // 原交易的metadata
	Mode                  string                 `json:"mode"`                    // sandbox 沙盒，live 生产
	Payer                 Payer                  `json:"payer"`                   // 付款人（玩家信息）
	ProjectId             int64                  `json:"project_id"`              // 应用id
	RefundId              string                 `json:"refund_id"`               // 退款id
	Status                string                 `json:"status"`                  // 退款状态
	TransactionId         string                 `json:"transaction_id"`          // 原交易id
	UpdateTime            int64                  `json:"update_time"`             // 更新时间
}

type Dispute struct {
	CreateTime            int64                  `json:"create_time"`    // 创建时间
	Amount                int64                  `json:"amount"`         // 金额
	Currency              string                 `json:"currency"`       // 币种
	DisputeId             string                 `json:"dispute_id"`     // 拒付id
	TransactionId         string                 `json:"transaction_id"` // 原始交易id
	Status                string                 `json:"status"`         // 拒付状态
	UpdateTime            int64                  `json:"update_time"`    // 更新时间
	Metadata              map[string]interface{} `json:"metadata"`       // 原交易的metadata
	MerchantTransactionId string                 `json:"merchant_transaction_id"`
	MerchantId            int64                  `json:"merchant_id"` // 商户id
	ProjectId             int64                  `json:"project_id"`  // 应用id
	Payer                 Payer                  `json:"payer"`       // 付款人（玩家信息）
	Mode                  string                 `json:"mode"`        // sandbox 沙盒，live 生产
}

type Names struct {
	LangId string `json:"lang_id"` // 多语言id
	EN     string `json:"en"`      // 英语
	CN     string `json:"cn"`      // 简中
	AR     string `json:"ar"`      // 阿拉伯语
	PL     string `json:"pl"`      // 波兰语
	DE     string `json:"de"`      // 德语
	RU     string `json:"ru"`      // 俄语
	FR     string `json:"fr"`      // 法语
	KO     string `json:"ko"`      // 韩语
	PT     string `json:"pt"`      // 葡萄牙语
	JA     string `json:"ja"`      // 日语
	TH     string `json:"th"`      // 泰语
	TR     string `json:"tr"`      // 土耳其语
	ES     string `json:"es"`      // 西班牙语
	ID     string `json:"id"`      // 印尼语
	IT     string `json:"it"`      // 意大利语
	VI     string `json:"vi"`      // 越南语
	TW     string `json:"tw"`      // 繁中
}

type Item struct {
	Amount         float64 `json:"amount"`          // 金额
	Code           string  `json:"code"`            // 商户端配置的道具id
	Currency       string  `json:"currency"`        // 币种
	Id             int64   `json:"id"`              // KOP商城中的道具id
	Names          Names   `json:"names"`           // 道具名称
	Num            int64   `json:"num"`             // 道具数量
	OrderId        int64   `json:"order_id"`        // 子订单id
	Price          float64 `json:"price"`           // 单价
	PurchaseLimits int64   `json:"purchase_limits"` // 限购数量
}

type GiftItem struct {
	GiftGoodsId string  `json:"gift_goods_id"` // 赠品id
	Logo        string  `json:"logo"`          // 赠品图标
	Num         float64 `json:"num"`           // 赠品数量，某些情况下，礼品数量可能出现小数
}

type SaleBiz struct {
	ActivityRule    string     `json:"activity_rule"`     // 活动规则
	ActivityType    int64      `json:"activity_type"`     // 活动类型
	ActivityTypeTxt string     `json:"activity_type_txt"` //
	BizId           int64      `json:"biz_id"`            // 活动id
	CanGift         bool       `json:"can_gift"`          // 是否赠送礼品
	GiftList        []GiftItem `json:"gift_list"`         // 赠品列表
	Name            string     `json:"name"`              // 活动名称
	ProjectId       int64      `json:"project_id"`        // 应用id
	StartTime       int64      `json:"start_time"`        // 活动开始时间
	StopTime        int64      `json:"stop_time"`         // 活动结束时间
}

type OrderMetadata struct {
	Gifts      []GiftItem `json:"gifts"`        // 赠品列表
	Items      []Item     `json:"items"`        // 商品（道具）列表
	OrderId    string     `json:"order_id"`     // 订单id
	SaleBizHit []SaleBiz  `json:"sale_biz_hit"` // 命中的活动
	Source     int64      `json:"source"`       // 订单来源，0、普通下单 1、购物车 2、批量购买商品
}

type Order struct {
	Amount          float64       `json:"amount"`            // 金额
	Country         string        `json:"country"`           // 国家
	CreateTime      int64         `json:"create_time"`       // 创建时间
	Currency        string        `json:"currency"`          // 币种
	Language        string        `json:"language"`          // 语种
	MerchantId      int64         `json:"merchant_id"`       // 商户id
	Metadata        OrderMetadata `json:"metadata"`          // 源数据，里面存储着订单信息、活动信息、赠品信息等
	Mode            string        `json:"mode"`              // 沙盒 sandbox, 生产 live
	PayId           string        `json:"pay_id"`            // 交易id
	PlayerId        string        `json:"player_id"`         // 玩家id
	ServerId        string        `json:"server_id"`         // 玩家服务器id
	PaymentType     int64         `json:"payment_type"`      // 订单类型
	ProjectId       int64         `json:"project_id"`        // 应用id
	Rfc1766Language string        `json:"rfc_1766_language"` // RPC1766语种
	Status          string        `json:"status"`            // 订单状态
	TransactionId   string        `json:"transaction_id"`    // 交易id
}

type BizInfo struct {
	Id        int64  `json:"id"`         // 活动id
	Name      string `json:"name"`       // 活动名称
	Type      int64  `json:"type"`       // 类型id
	StartTime int64  `json:"start_time"` // 活动开始时间
	StopTime  int64  `json:"Stop_time"`  // 活动结束时间
	BizType   string `json:"biz_type"`   // 活动类型
}

type GiftInfo struct {
	GiftList []GiftItem `json:"gift_list"` // 赠品列表
}

type Gift struct {
	MerchantId      int64                  `json:"merchant_id"`       // 商户id
	ProjectId       int64                  `json:"project_id"`        // 应用id
	PlayerId        string                 `json:"player_id"`         // 玩家id
	ServerId        string                 `json:"server_id"`         // 玩家服务器id
	Mode            string                 `json:"mode"`              // 模式
	BizInfo         BizInfo                `json:"biz_info"`          // 活动信息
	GiftInfo        GiftInfo               `json:"gift_info"`         // 赠品信息
	CreateTime      int64                  `json:"create_time"`       // 创建时间
	Language        string                 `json:"language"`          // 语种
	Rfc1766Language string                 `json:"rfc_1766_language"` // rfc1766语种
	Metadata        map[string]interface{} `json:"metadata"`          // metadata
}

type NotificationUser struct {
	Id         string `json:"id"`          // 玩家id
	Server     string `json:"server"`      // 玩家服务器id
	MerchantId int64  `json:"merchant_id"` // 商户id
	ProjectId  int64  `json:"project_id"`  // 应用id
}

type Notification struct {
	EventId          string                 `json:"event_id"`          // 消息id
	MerchantId       int64                  `json:"merchant_id"`       // 商户id
	ProjectId        int64                  `json:"project_id"`        // 应用id
	NotificationType string                 `json:"notification_type"` // 消息类型
	User             NotificationUser       `json:"user"`              // 查询玩家信息时，需要用到
	Data             map[string]interface{} `json:"data"`              // 消息数据
}
