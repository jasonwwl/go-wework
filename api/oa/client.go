package oa

import "github.com/jasonwwl/go-wework"

type OAClient struct {
	client *wework.Client
}

func NewOAClient(client *wework.Client) *OAClient {
	return &OAClient{client: client}
}
