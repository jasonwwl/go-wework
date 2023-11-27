package wework_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jasonwwl/go-wework"
)

func TestBuildKey(t *testing.T) {
	client := newTestClient()
	openCorpCfg := client.GetOpenCorpConfig()
	internalCorpCfg := client.GetInternalCorpConfig()
	tests := []struct {
		tokenType wework.TokenType
		want      string
	}{
		{
			tokenType: wework.AccessToken.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.AccessToken.TokenType, internalCorpCfg.CorpID),
		},
		{
			tokenType: wework.ProviderToken.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.ProviderToken.TokenType, openCorpCfg.ProviderCorpID),
		},
		{
			tokenType: wework.SuiteToken.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.SuiteToken.TokenType, openCorpCfg.SuiteID),
		},
		{
			tokenType: wework.AuthCorpAccessToken.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.AuthCorpAccessToken.TokenType, openCorpCfg.AuthCorpID),
		},
		{
			tokenType: wework.PermanentCode.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.PermanentCode.TokenType, openCorpCfg.AuthCorpID),
		},
		{
			tokenType: wework.SuiteTicket.TokenType,
			want:      fmt.Sprintf("%s:%s", wework.SuiteTicket.TokenType, openCorpCfg.SuiteID),
		},
	}

	for _, test := range tests {
		got, err := wework.BuildKey(client, test.tokenType)

		if err != nil {
			t.Errorf("BuildKey returned an error: %v", err)
		}

		if got != test.want {
			t.Errorf("BuildKey returned unexpected key: got %v want %v", got, test.want)
		}
	}
}

func TestSetToken(t *testing.T) {
	store := wework.NewMemoryStore()
	client := newTestClient()
	ctx := context.TODO()

	tokenType := wework.AccessToken.TokenType
	token := "testToken"
	expiresIn := int64(60)

	err := store.SetToken(client, ctx, tokenType, token, expiresIn)

	if err != nil {
		t.Errorf("SetToken returned an error: %v", err)
	}
}

func TestGetToken(t *testing.T) {
	store := wework.NewMemoryStore()
	client := newTestClient()
	ctx := context.TODO()

	tokenType := wework.AccessToken.TokenType
	token := "testToken"
	expiresIn := int64(1) // 5 seconds

	// 首先设置一个令牌
	err := store.SetToken(client, ctx, tokenType, token, expiresIn)
	if err != nil {
		t.Fatalf("SetToken returned an error: %v", err)
	}

	// 尝试获取同一个令牌
	retrievedToken, err := store.GetToken(client, ctx, tokenType)
	if err != nil {
		t.Fatalf("GetToken returned an error: %v", err)
	}

	if retrievedToken != token {
		t.Errorf("GetToken returned unexpected token: got %v want %v", retrievedToken, token)
	}

	// 测试过期的令牌
	time.Sleep(time.Second * 1) // 等待超过令牌的过期时间
	_, err = store.GetToken(client, ctx, tokenType)
	if err == nil {
		t.Errorf("Expected an error for expired token, but got none")
	}
}
