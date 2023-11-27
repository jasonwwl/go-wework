package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jasonwwl/go-wework"
	"github.com/jasonwwl/go-wework/api/basic"
)

func main() {
	clientCfg, err := wework.NewConfigBuilder().InternalCorp(wework.InternalCorp{
		CorpID:  os.Getenv("WEWORK_CORP_ID"),
		AgentID: os.Getenv("WEWORK_AGENT_ID"),
		Secret:  os.Getenv("WEWORK_SECRET"),
	}).Build()
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
