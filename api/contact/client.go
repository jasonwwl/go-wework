package contact

import "github.com/jasonwwl/go-wework"

type ContactClient struct {
	client *wework.Client
}

func NewContactClient(client *wework.Client) *ContactClient {
	return &ContactClient{client: client}
}
