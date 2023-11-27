package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jasonwwl/go-wework"
	"github.com/jasonwwl/go-wework/api/basic"
)

func main() {
	clientCfg, err := wework.NewConfigBuilder().OpenCorp(wework.OpenCorp{
		ProviderCorpID:   os.Getenv("WEWORK_PROVIDER_CORP_ID"),
		ProviderSecret:   os.Getenv("WEWORK_PROVIDER_SECRET"),
		SuiteID:          os.Getenv("WEWORK_SUITE_ID"),
		SuiteSecret:      os.Getenv("WEWORK_SUITE_SECRET"),
		SuiteToken:       os.Getenv("WEWORK_SUITE_TOKEN"),
		SuiteEncodingAES: os.Getenv("WEWORK_SUITE_AES_KEY"),
		AuthCorpID:       os.Getenv("WEWORK_AUTH_CORP_ID"),
	}).Build()
	if err != nil {
		fmt.Printf("Build open corp config error: %v\n", err)
		return
	}

	client := wework.NewClient(clientCfg)

	store := client.GetStore()

	// 存入最新的suite ticket
	// 此逻辑应该在回调接口中实现
	err = store.SetToken(client, context.Background(), wework.SuiteTicket.TokenType, "ticket...", 600)
	if err != nil {
		fmt.Printf("SetSuiteTicket error: %v\n", err)
		return
	}

	// 存入授权企业的永久授权码
	// 此逻辑应该在企业授权完成后实现
	err = store.SetToken(client, context.Background(), wework.PermanentCode.TokenType, "code...", 0)
	if err != nil {
		fmt.Printf("SetPermanentCode error: %v\n", err)
		return
	}

	basicClient := basic.NewWithClient(client)

	resp, err := basicClient.GetAdminList(context.Background(), client.GetOpenCorpConfig().AuthCorpID, 100001)

	if err != nil {
		fmt.Printf("GetAdminList error: %v\n", err)
		return
	}

	fmt.Printf("Admin: %v\n", resp.Admin)
}
