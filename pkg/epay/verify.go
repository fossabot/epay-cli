package epay

import (
	"context"
	"github.com/mitchellh/mapstructure"
)

type VerifyRes struct {
	Type           PaymentType // 支付类型
	TradeNo        string      `mapstructure:"trade_no"`     // 易支付订单号
	ServiceTradeNo string      `mapstructure:"out_trade_no"` // 商家订单号
	Name           string      ``                            // 商品名称
	Money          string      ``                            // 金额
	TradeStatus    string      `mapstructure:"trade_status"` // 订单支付状态
	VerifyStatus   bool        `mapstructure:"-"`            // 签名检验
}

func (c *Client) Verify(_ context.Context, params map[string]string) (*VerifyRes, error) {
	sign := params["sign"]
	var verifyRes VerifyRes
	err := mapstructure.Decode(params, &verifyRes)

	verifyRes.VerifyStatus = sign == GenerateParams(params, c.config.AppSecret)["sign"]
	if err != nil {
		return nil, err
	}

	return &verifyRes, nil
}
