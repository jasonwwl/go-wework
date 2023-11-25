package wework_test

import (
	"context"
	"testing"
	"time"

	"github.com/jasonwwl/go-wework"
)

func TestFetchAccessToken(t *testing.T) {
	client := newTestClient()
	startTime := time.Now()
	accessToken, err := client.FetchAccessToken(context.TODO())
	duration := time.Since(startTime)
	t.Logf("FetchAccessToken took %s to complete", duration)

	if err != nil {
		t.Errorf("FetchAccessToken returned an error: %v", err)
	}

	if accessToken == "" {
		t.Errorf("FetchAccessToken returned empty token")
	}

	t.Logf("accessToken: %s", accessToken)

	sAk, err := client.GetStore().GetToken(client, wework.AccessToken.TokenType)

	if err != nil {
		t.Errorf("GetToken returned an error: %v", err)
	}

	if sAk != accessToken {
		t.Errorf("GetToken returned unexpected token: got %v want %v", sAk, accessToken)
	}
}

func TestFetchSuiteToken(t *testing.T) {
	client := newTestClient()
	client.InitStoreData([]wework.StoreData{
		{
			TokenDescriptor: wework.SuiteTicket,
			Token:           "l0L8tMm7-92yIUw9N-Snp7Ks2EhQYTfbwep6pnlCqYDKcgOVhjy5pJHjDqDuYGpr",
		},
	})

	suiteToken, err := client.FetchSuiteToken(context.TODO())
	if err != nil {
		t.Errorf("FetchSuiteToken returned an error: %v", err)
	}

	if suiteToken == "" {
		t.Errorf("FetchSuiteToken returned empty token")
	}

	t.Logf("suiteToken: %s", suiteToken)

	sAk, err := client.GetStore().GetToken(client, wework.SuiteToken.TokenType)

	if err != nil {
		t.Errorf("GetToken returned an error: %v", err)
	}

	if sAk != suiteToken {
		t.Errorf("GetToken returned unexpected token: got %v want %v", sAk, suiteToken)
	}
}
