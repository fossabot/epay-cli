package epay

import (
	"context"
	"net/url"
	"strconv"
)

type SubmitArgs struct {
	Type           PaymentType // 支付类型
	ServiceTradeNo string      // 商家订单号
	Name           string      // 商品名称
	Money          string      // 金额
	NotifyUrl      string      // 异步通知地址
	ReturnUrl      string      // 返回地址
}

// Submit implements Service.Submit
func (c *Client) Submit(_ context.Context, args *SubmitArgs) (string, map[string]string, error) {
	requestParams := map[string]string{
		"pid":          strconv.Itoa(c.config.PartnerID),
		"type":         string(args.Type),
		"out_trade_no": args.ServiceTradeNo,
		"notify_url":   args.NotifyUrl,
		"name":         args.Name,
		"money":        args.Money,
		"sign_type":    "MD5",
		"return_url":   args.ReturnUrl,
		"sign":         "",
	}

	u, err := url.Parse(c.reqClient.BaseURL)
	if err != nil {
		return "", nil, err
	}

	u, err = u.Parse(SubmitUrl)
	if err != nil {
		return "", nil, err
	}

	return u.String(), GenerateParams(requestParams, c.config.AppSecret), nil
}
