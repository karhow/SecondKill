package global

var (
	// PushToken 集成pushlus 一键免费推送消息 令牌
	// http://www.pushplus.plus/doc/guide/api.html#%E4%B8%80%E3%80%81%E5%8F%91%E9%80%81%E6%B6%88%E6%81%AF%E6%8E%A5%E5%8F%A3
	PushToken = ""
	// PushToken      = "xx"
	PushplusReqURL = "http://www.pushplus.plus/send"
)

// PushplusReq pushplus 消息发送请求体
// 请求地址：http://www.pushplus.plus/send
type PushplusReq struct {
	// Token
	// Required
	Token string `json:"token"`
	Title string `json:"title"`
	// Content
	// Required
	Content     string `json:"content"`
	Topic       string `json:"topic"`
	Template    string `json:"template"`
	Channel     string `json:"channel"`
	Webhook     string `json:"webhook"`
	CallbackUrl string `json:"callbackUrl"`
	Timestamp   string `json:"timestamp"`
}
