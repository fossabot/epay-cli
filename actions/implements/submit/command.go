package submit

import (
	"github.com/AH-dark/epay-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

func (svc *service) Command() *cli.Command {
	return &cli.Command{
		Name:   "submit",
		Usage:  "get submit.php url and args",
		Action: svc.Do,
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "pid", Required: true, Usage: "partner id", EnvVars: []string{"EPAY_PID"}},
			&cli.StringFlag{Name: "secret", Required: true, Usage: "app secret", EnvVars: []string{"EPAY_SECRET"}},
			&cli.StringFlag{Name: "endpoint", Required: true, Usage: "endpoint", EnvVars: []string{"EPAY_ENDPOINT"}},
			&cli.StringFlag{Name: "type", Value: "alipay", Usage: "payment type", Required: false, EnvVars: []string{"EPAY_TYPE"}},
			&cli.StringFlag{Name: "trade-no", Usage: "service trade no", Value: utils.RandString(32), Required: false, EnvVars: []string{"EPAY_TRADE_NO"}},
			&cli.StringFlag{Name: "name", Usage: "name", Value: "cli 测试商品", Required: false, EnvVars: []string{"EPAY_NAME"}},
			&cli.StringFlag{Name: "money", Usage: "money", Value: "1.00", Required: false, EnvVars: []string{"EPAY_MONEY"}},
			&cli.StringFlag{Name: "notify-url", Usage: "notify url", Required: false, EnvVars: []string{"EPAY_NOTIFY_URL"}},
			&cli.StringFlag{Name: "return-url", Usage: "return url", Required: false, EnvVars: []string{"EPAY_RETURN_URL"}},
		},
	}
}
