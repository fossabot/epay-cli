package epay

type PaymentType string

var (
	Alipay    PaymentType = "alipay" // Alipay 支付宝
	WechatPay PaymentType = "wxpay"  // WechatPay 微信
)

type DeviceType string

var (
	PC     DeviceType = "pc"     // PC PC端
	MOBILE DeviceType = "mobile" // MOBILE 移动端
	WECHAT DeviceType = "wechat" // WECHAT 微信
)

type CommonErrorRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
