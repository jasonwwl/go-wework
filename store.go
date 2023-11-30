package wework

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Store interface {
	GetToken(client *Client, ctx context.Context, tokenType TokenType) (string, error)
	SetToken(client *Client, ctx context.Context, tokenType TokenType, token string, expiresIn time.Duration) error
}

type TokenInfo struct {
	token     string
	expiresAt time.Time
}

type MemoryStore struct {
	tokens sync.Map
}

var _ Store = (*MemoryStore)(nil)

var (
	instance Store
	once     sync.Once
)

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{tokens: sync.Map{}}
}

func InitMemoryStore() Store {
	once.Do(func() {
		instance = NewMemoryStore()
	})
	return instance
}

func (m *MemoryStore) GetToken(client *Client, ctx context.Context, tokenType TokenType) (string, error) {
	key, err := BuildKey(client, tokenType)
	if err != nil {
		return "", err
	}

	val, ok := m.tokens.Load(key)
	if !ok {
		return "", ErrNilStoreToken
	}

	tokenInfo, ok := val.(*TokenInfo)
	if !ok {
		return "", fmt.Errorf("invalid token info type")
	}

	if !tokenInfo.expiresAt.IsZero() && tokenInfo.expiresAt.Before(time.Now()) {
		return "", ErrNilStoreToken
	}

	return tokenInfo.token, nil

}

func (m *MemoryStore) SetToken(client *Client, ctx context.Context, tokenType TokenType, token string, expiresIn time.Duration) error {
	key, err := BuildKey(client, tokenType)
	if err != nil {
		return err
	}

	m.tokens.Store(key, MakeTokenInfo(token, expiresIn))
	return nil
}

func MakeTokenInfo(token string, expiresIn time.Duration) *TokenInfo {
	// 如果exipresIn为0，则表示永久有效

	var expiresAt time.Time
	if expiresIn > 0 {
		expiresAt = time.Now().Add(expiresIn)
	} else {
		expiresAt = time.Time{}
	}

	return &TokenInfo{token: token, expiresAt: expiresAt}
}

func BuildKey(client *Client, tokenType TokenType) (string, error) {
	var id string

	switch tokenType {

	case ProviderToken.TokenType:
		id = client.config.OpenCorp.ProviderCorpID

	case SuiteToken.TokenType, SuiteTicket.TokenType:
		id = client.config.OpenCorp.SuiteID

	case AccessToken.TokenType:
		id = client.config.InternalCorp.CorpID

	case AuthCorpAccessToken.TokenType, PermanentCode.TokenType:
		id = client.config.OpenCorp.SuiteID + ":" + client.config.OpenCorp.AuthCorpID

	default:
		return "", fmt.Errorf("invalid token type: %s", tokenType)
	}

	if id == "" {
		return "", fmt.Errorf("id is empty for token type: %s", tokenType)
	}

	return fmt.Sprintf("%s:%s", tokenType, id), nil
}
