package submit

import (
	"fmt"
	"net/url"

	"github.com/urfave/cli/v2"

	"github.com/AH-dark/epay-cli/actions/factory"
	"github.com/AH-dark/epay-cli/pkg/epay"
)

type service struct {
}

func NewService() factory.ActionService {
	return &service{}
}

func (svc *service) Do(c *cli.Context) error {
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
