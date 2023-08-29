package actions

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/AH-dark/epay-cli/pkg/epay"
	"github.com/AH-dark/epay-cli/pkg/utils"
)

func MapiCommand() *cli.Command {
	return &cli.Command{
		Name:   "mapi",
		Usage:  "mapi submit",
		Action: MapiAction,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "pid",
				Required: true,
				Usage:    "partner id",
				EnvVars:  []string{"EPAY_PID"},
			},
			&cli.StringFlag{
				Name:     "secret",
				Required: true,
				Usage:    "app secret",
				EnvVars:  []string{"EPAY_SECRET"},
			},
			&cli.StringFlag{
				Name:     "endpoint",
				Required: true,
				Usage:    "endpoint",
				EnvVars:  []string{"EPAY_ENDPOINT"},
			},
			&cli.StringFlag{
				Name:     "type",
				Value:    "alipay",
				Usage:    "payment type",
				Required: false,
				EnvVars:  []string{"EPAY_TYPE"},
			},
			&cli.StringFlag{
				Name:     "trade-no",
				Usage:    "service trade no",
				Value:    utils.RandString(32),
				Required: false,
				EnvVars:  []string{"EPAY_TRADE_NO"},
			},
			&cli.StringFlag{
				Name:     "name",
				Usage:    "name",
				Value:    "cli 测试商品",
				Required: false,
				EnvVars:  []string{"EPAY_NAME"},
			},
			&cli.StringFlag{
				Name:     "money",
				Usage:    "money",
				Value:    "1.00",
				Required: false,
				EnvVars:  []string{"EPAY_MONEY"},
			},
			&cli.StringFlag{
				Name:     "notify-url",
				Usage:    "notify url",
				Required: false,
				EnvVars:  []string{"EPAY_NOTIFY_URL"},
			},
			&cli.StringFlag{
				Name:     "return-url",
				Usage:    "return url",
				Required: false,
				EnvVars:  []string{"EPAY_RETURN_URL"},
			},
			&cli.StringFlag{
				Name:     "client-ip",
				Usage:    "client ip",
				Value:    "127.0.0.1",
				Required: false,
				EnvVars:  []string{"EPAY_CLIENT_IP"},
			},
			&cli.StringFlag{
				Name:     "device",
				Usage:    "device",
				Value:    "pc",
				Required: false,
				EnvVars:  []string{"EPAY_DEVICE"},
			},
			&cli.StringFlag{
				Name:     "param",
				Usage:    "param",
				Required: false,
				EnvVars:  []string{"EPAY_PARAM"},
			},
		},
	}
}

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
