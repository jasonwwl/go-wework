package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jasonwwl/go-wework"
	"github.com/jasonwwl/go-wework/api/basic"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisStore := NewRedisStore(rdb)

	clientCfg, err := wework.NewConfigBuilder().InternalCorp(
		wework.InternalCorp{
			CorpID:  os.Getenv("WEWORK_CORP_ID"),
			AgentID: os.Getenv("WEWORK_AGENT_ID"),
			Secret:  os.Getenv("WEWORK_SECRET"),
		},
	).TokenStore(redisStore).Build()
	if err != nil {
		fmt.Printf("Build internal corp config error: %v\n", err)
		return
	}

	client := wework.NewClient(clientCfg)

	basicClient := basic.NewWithClient(client)

	resp, err := basicClient.UserIDToOpenUserID(context.Background(), []string{"jasonwwl"})
	if err != nil {
		fmt.Printf("UserIDToOpenUserID error: %v\n", err)
		return
	}

	fmt.Printf("OpenUserIDList: %v, InvalidUserIDList: %v\n", resp.OpenUserIDList, resp.InvalidUserIDList)
}
