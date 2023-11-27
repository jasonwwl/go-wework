package main

import (
	"context"
	"time"

	"github.com/jasonwwl/go-wework"
	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{rdb: rdb}
}

func (s *RedisStore) GetToken(client *wework.Client, ctx context.Context, tokenType wework.TokenType) (string, error) {
	key, err := wework.BuildKey(client, tokenType)
	if err != nil {
		return "", err
	}

	res, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	if res == "" {
		return "", wework.ErrNilStoreToken
	}

	return res, nil
}

func (s *RedisStore) SetToken(client *wework.Client, ctx context.Context, tokenType wework.TokenType, token string, expiresIn int64) error {
	key, err := wework.BuildKey(client, tokenType)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, key, token, time.Second*time.Duration(expiresIn)).Err()
}
