package auth

import (
	"github.com/jasonwwl/go-wework"
)

type AuthClient struct {
	client *wework.Client
}

func New(client *wework.Client) *AuthClient {
	return &AuthClient{client: client}
}
