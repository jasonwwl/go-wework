package basic

import "github.com/jasonwwl/go-wework"

type BasicClient struct {
	client *wework.Client
}

func NewBasicClient(client *wework.Client) *BasicClient {
	return &BasicClient{client: client}
}
