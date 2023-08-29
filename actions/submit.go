package actions

import (
	"fmt"
	"net/url"

	"github.com/AH-dark/epay-cli/pkg/epay"
	"github.com/AH-dark/epay-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

func SubmitCommand() *cli.Command {
	return &cli.Command{
		Name:   "submit",
		Usage:  "get submit.php url and args",
		Action: SubmitAction,
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
		},
	}
}

func SubmitAction(c *cli.Context) error {
	client, err := epay.NewClient(&epay.Config{
		PartnerID: c.Int("pid"),
		AppSecret: c.String("secret"),
		Endpoint:  c.String("endpoint"),
	})
	if err != nil {
		return err
	}

	u, args, err := client.Submit(c.Context, &epay.SubmitArgs{
		Type:           epay.PaymentType(c.String("type")),
		ServiceTradeNo: c.String("trade-no"),
		Name:           c.String("name"),
		Money:          c.String("money"),
		NotifyUrl:      c.String("notify-url"),
		ReturnUrl:      c.String("return-url"),
	})
	if err != nil {
		return err
	}

	values := url.Values{}
	for k, v := range args {
		values.Add(k, v)
	}

	urlObj, err := url.Parse(u)
	if err != nil {
		return err
	}

	urlObj.RawQuery = values.Encode()

	fmt.Println(urlObj.String())

	return nil
}
