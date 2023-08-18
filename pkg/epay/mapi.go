package epay

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AH-dark/epay-cli/pkg/utils"
)

type MApiSubmitArgs struct {
	Type       PaymentType `json:"type"`             // Type 支付类型
	OutTradeNo string      `json:"out_trade_no"`     // OutTradeNo 商户订单号
	NotifyUrl  string      `json:"notify_url"`       // NotifyUrl 异步通知地址
	ReturnUrl  *string     `json:"return_url"`       // ReturnUrl 同步通知地址
	Name       string      `json:"name"`             // Name 商品名称
	Money      string      `json:"money"`            // Money 金额 保留两位小数
	ClientIP   string      `json:"clientip"`         // ClientIP 客户端IP
	Device     *DeviceType `json:"device,omitempty"` // Device 设备类型
	Param      *string     `json:"param,omitempty"`  // Param 附加参数
	Sign       string      `json:"sign"`             // Sign 签名
	SignType   string      `json:"sign_type"`        // SignType 签名类型
}

type MApiSubmitRes struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	TradeNo   string `json:"trade_no"`
	PayUrl    string `json:"payurl"`
	QRCode    string `json:"qrcode"`
	UrlScheme string `json:"urlscheme"`
}

func (c *Client) MApiSubmit(ctx context.Context, args *MApiSubmitArgs) (*MApiSubmitRes, *http.Response, error) {
	requestArgs := GenerateParams(map[string]string{
		"pid":          strconv.Itoa(c.config.PartnerID),
		"type":         string(args.Type),
		"out_trade_no": args.OutTradeNo,
		"notify_url":   args.NotifyUrl,
		"return_url":   utils.ParseEmptyPtr(args.ReturnUrl),
		"name":         args.Name,
		"money":        args.Money,
		"clientip":     args.ClientIP,
		"device":       string(utils.ParseEmptyPtr(args.Device)),
		"param":        utils.ParseEmptyPtr(args.Param),
		"sign_type":    "MD5",
		"sign":         "",
	}, c.config.AppSecret)

	var res MApiSubmitRes
	resp, err := c.reqClient.R().
		SetContext(ctx).
		SetFormData(requestArgs).
		SetSuccessResult(&res).
		Post(MAPIUrl)
	if err != nil {
		return nil, resp.Response, err
	}

	if !resp.IsSuccessState() || resp.StatusCode != http.StatusOK {
		if errRes, ok := resp.ErrorResult().(CommonErrorRes); ok {
			return nil, resp.Response, errors.New(errRes.Msg)
		}

		return nil, resp.Response, fmt.Errorf("bad response status: %d", resp.StatusCode)
	}

	if res.Code != 1 {
		return nil, resp.Response, fmt.Errorf("epay error: %s", res.Msg)
	}

	return &res, resp.Response, nil
}
