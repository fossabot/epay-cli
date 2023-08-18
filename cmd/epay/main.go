package main

import (
	"log"
	"os"

	"github.com/AH-dark/epay-cli/actions"
	"github.com/AH-dark/epay-cli/pkg/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "epay-cli",
		Version: "0.0.1",
		Usage:   "epay-cli is a command line tool for epay",
		Commands: []*cli.Command{
			{
				Name:     "submit",
				Usage:    "get submit.php url and args",
				Category: "test",
				Action:   actions.SubmitAction,
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "pid",
						Required: true,
						Usage:    "partner id",
					},
					&cli.StringFlag{
						Name:     "secret",
						Required: true,
						Usage:    "app secret",
					},
					&cli.StringFlag{
						Name:     "endpoint",
						Required: true,
						Usage:    "endpoint",
					},
					&cli.StringFlag{
						Name:     "type",
						Value:    "alipay",
						Usage:    "payment type",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "trade-no",
						Usage:    "service trade no",
						Value:    utils.RandString(32),
						Required: false,
					},
					&cli.StringFlag{
						Name:     "name",
						Usage:    "name",
						Value:    "cli 测试商品",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "money",
						Usage:    "money",
						Value:    "1.00",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "notify-url",
						Usage:    "notify url",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "return-url",
						Usage:    "return url",
						Required: false,
					},
				},
			},
			{
				Name:     "mapi",
				Usage:    "mapi submit",
				Category: "test",
				Action:   actions.MapiAction,
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:     "pid",
						Required: true,
						Usage:    "partner id",
					},
					&cli.StringFlag{
						Name:     "secret",
						Required: true,
						Usage:    "app secret",
					},
					&cli.StringFlag{
						Name:     "endpoint",
						Required: true,
						Usage:    "endpoint",
					},
					&cli.StringFlag{
						Name:     "type",
						Value:    "alipay",
						Usage:    "payment type",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "trade-no",
						Usage:    "service trade no",
						Value:    utils.RandString(32),
						Required: false,
					},
					&cli.StringFlag{
						Name:     "name",
						Usage:    "name",
						Value:    "cli 测试商品",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "money",
						Usage:    "money",
						Value:    "1.00",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "notify-url",
						Usage:    "notify url",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "return-url",
						Usage:    "return url",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "client-ip",
						Usage:    "client ip",
						Value:    "127.0.0.1",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "device",
						Usage:    "device",
						Value:    "pc",
						Required: false,
					},
					&cli.StringFlag{
						Name:     "param",
						Usage:    "param",
						Required: false,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("error: ", err)
	}
}
