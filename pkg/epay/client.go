package epay

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/imroc/req/v3"
)

const (
	SubmitUrl = "/submit.php"
	MAPIUrl   = "/mapi.php"
)

type Config struct {
	PartnerID int
	AppSecret string
	Endpoint  string
}

var _ Service = (*Client)(nil)

// Service 易支付API
type Service interface {
	Submit(ctx context.Context, args *SubmitArgs) (string, map[string]string, error) // Submit 生成支付链接和参数
	Verify(ctx context.Context, params map[string]string) (*VerifyRes, error)        // Verify 验证回调参数是否符合签名

	MApiSubmit(ctx context.Context, args *MApiSubmitArgs) (*MApiSubmitRes, *http.Response, error) // MApiSubmit 生成支付链接和参数
}

type Client struct {
	config    *Config
	reqClient *req.Client
}

// NewClient 创建一个新的易支付客户端
func NewClient(config *Config) (*Client, error) {
	if _, err := url.Parse(config.Endpoint); err != nil {
		return nil, err
	}

	client := req.NewClient()
	client.SetBaseURL(config.Endpoint)
	client.SetTimeout(time.Second * 10)
	client.SetCommonErrorResult(&CommonErrorRes{})

	return &Client{
		config:    config,
		reqClient: client,
	}, nil
}
