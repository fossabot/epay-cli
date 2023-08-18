package actions

import (
	"fmt"

	"github.com/AH-dark/epay-cli/pkg/epay"
	"github.com/urfave/cli/v2"
)

func SubmitAction(c *cli.Context) error {
	client, err := epay.NewClient(&epay.Config{
		PartnerID: c.Int("pid"),
		AppSecret: c.String("secret"),
		Endpoint:  c.String("endpoint"),
	})
	if err != nil {
		return err
	}

	url, args, err := client.Submit(c.Context, &epay.SubmitArgs{
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

	fmt.Println("URL:", url)
	fmt.Println("Args:", args)

	return nil
}
