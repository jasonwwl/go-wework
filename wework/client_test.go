package wework_test

import (
	"bufio"
	"os"
	"strings"

	"github.com/jasonwwl/go-wework/wework"
)

func LoadEnv() error {
	file, err := os.Open("../.env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		os.Setenv(parts[0], parts[1])
	}

	return scanner.Err()
}

func newTestClient() *wework.Client {
	LoadEnv()

	cfg, err := wework.NewConfigBuilder().InternalCorp(wework.InternalCorp{
		CorpID:  os.Getenv("WEWORK_CORP_ID"),
		Secret:  os.Getenv("WEWORK_SECRET"),
		AgentID: os.Getenv("WEWORK_AGENT_ID"),
	}).OpenCorp(wework.OpenCorp{
		ProviderCorpID:   os.Getenv("WEWORK_PROVIDER_CORP_ID"),
		ProviderSecret:   os.Getenv("WEWORK_PROVIDER_SECRET"),
		SuiteID:          os.Getenv("WEWORK_SUITE_ID"),
		SuiteSecret:      os.Getenv("WEWORK_SUITE_SECRET"),
		SuiteToken:       os.Getenv("WEWORK_SUITE_TOKEN"),
		SuiteEncodingAES: os.Getenv("WEWORK_SUITE_AES_KEY"),
		AuthCorpID:       os.Getenv("WEWORK_AUTH_CORP_ID"),
	}).Build()

	if err != nil {
		panic(err)
	}

	return wework.NewClient(cfg)
}
