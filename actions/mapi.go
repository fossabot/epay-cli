package actions

import (
	"fmt"
	"github.com/AH-dark/epay-cli/pkg/utils"
	"strconv"

	"github.com/AH-dark/epay-cli/pkg/epay"
	"github.com/urfave/cli/v2"
)

func MapiAction(c *cli.Context) error {
	client, err := epay.NewClient(&epay.Config{
		PartnerID: c.Int("pid"),
		AppSecret: c.String("secret"),
		Endpoint:  c.String("endpoint"),
	})
	if err != nil {
		return err
	}

	sign := utils.CalculateEPaySign(map[string]string{
		"pid":          strconv.Itoa(c.Int("pid")),
		"type":         c.String("type"),
		"out_trade_no": c.String("trade-no"),
		"notify_url":   c.String("notify-url"),
		"return_url":   c.String("return-url"),
		"name":         c.String("name"),
		"money":        c.String("money"),
		"clientip":     c.String("client-ip"),
		"device":       c.String("device"),
		"param":        c.String("param"),
	}, c.String("secret"))
	fmt.Println("Sign:", sign)

	url, args, err := client.MApiSubmit(c.Context, &epay.MApiSubmitArgs{
		Type:       epay.PaymentType(c.String("type")),
		OutTradeNo: c.String("trade-no"),
		Name:       c.String("name"),
		Money:      c.String("money"),
		NotifyUrl:  c.String("notify-url"),
		ReturnUrl:  utils.EmptyPtr(c.String("return-url")),
		ClientIP:   c.String("client-ip"),
		Device:     utils.EmptyPtr(epay.DeviceType(c.String("device"))),
		Param:      utils.EmptyPtr(c.String("param")),
		Sign:       sign,
		SignType:   "MD5",
	})
	if err != nil {
		return err
	}

	fmt.Println("URL:", url)
	fmt.Println("Args:", args)

	return nil
}
