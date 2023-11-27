package basic

import (
	"github.com/jasonwwl/go-wework"
)

type BasicClient struct {
	client *wework.Client
}

func NewWithClient(client *wework.Client) *BasicClient {
	return &BasicClient{client: client}
}

func NewWithConfig(cfg *wework.ClientConfig) *BasicClient {
	return &BasicClient{client: wework.NewClient(cfg)}
}
