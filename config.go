package wework

import (
	"context"
	"net/http"
	"time"
)

const (
	WeWorkAPIURL = "https://qyapi.weixin.qq.com/cgi-bin"
	Timeout      = time.Second * 30
)

type TokenType string

type TokenDescriptor struct {
	TokenType  TokenType
	ParamValue string
}

var (
	AccessToken         = &TokenDescriptor{TokenType: "AccessToken", ParamValue: "access_token"}
	ProviderToken       = &TokenDescriptor{TokenType: "ProviderAccessToken", ParamValue: "provider_access_token"}
	SuiteToken          = &TokenDescriptor{TokenType: "SuiteAccessToken", ParamValue: "suite_access_token"}
	AuthCorpAccessToken = &TokenDescriptor{TokenType: "AuthCorpAccessToken", ParamValue: "access_token"} // 注意这里的区别
	SuiteTicket         = &TokenDescriptor{TokenType: "SuiteTicket", ParamValue: "suite_ticket"}
	PermanentCode       = &TokenDescriptor{TokenType: "PermanentCode", ParamValue: "permanent_code"}
)

type ClientConfig struct {
	Options      *Options
	InternalCorp *InternalCorp
	OpenCorp     *OpenCorp
	HTTPClient   *http.Client
}

type InternalCorp struct {
	CorpID  string
	Secret  string
	AgentID string
}

type OpenCorp struct {
	ProviderCorpID   string
	ProviderSecret   string
	SuiteID          string
	SuiteSecret      string
	SuiteToken       string
	SuiteEncodingAES string
	AuthCorpID       string
}

type Options struct {
	Timeout    time.Duration
	TokenStore Store
	BaseURL    string
}

func (c *Client) GetStore() Store {
	return c.config.Options.TokenStore
}

func (c *Client) GetConfig() *ClientConfig {
	return c.config
}

func (c *Client) GetInternalCorpConfig() *InternalCorp {
	return c.GetConfig().InternalCorp
}

func (c *Client) GetOpenCorpConfig() *OpenCorp {
	return c.GetConfig().OpenCorp
}

type StoreData struct {
	TokenDescriptor *TokenDescriptor
	Token           string
}

func (c *Client) InitStoreData(data []StoreData, ctx context.Context) error {
	if ctx == nil {
		ctx = context.TODO()
	}

	for _, item := range data {
		err := c.GetStore().SetToken(c, ctx, item.TokenDescriptor.TokenType, item.Token, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) SetAuthCorpID(authCorpID string) {
	c.GetOpenCorpConfig().AuthCorpID = authCorpID
}
