# Go WeWork

This library provides unofficial Go clients for [WeWork API](https://developer.work.weixin.qq.com/document). We support:

- InternalCorpAPI
- OpenCorpAPI (SaaS)
- Encrypt/Decrypt WeWork event messages

## Installation

```shell
go get github.com/jasonwwl/go-wework
```

Currently, go-wework requires Go version 1.18 or greater.

## Usage

### InternalCorpAPI example usage

```go
package main

func main() {
    cfg, err := wework.NewConfigBuilder().InternalCorp(wework.InternalCorp{
        CorpID:  "your corpid",
        Secret:  "your app secret",
        AgentID: "your app agentid"
    }).Build()

    if err != nil {
        fmt.Printf("InternalCorp config build error: %v\n", err)
        return
    }

    client := wework.NewClient(cfg)

    resp, err := client.GetExternalContactList("zhangsan")
    if err != nil {
        fmt.Printf("GetExternalContactList error: %v\n", err)
        return
    }

    fmt.Printf("GetExternalContactList result: %v\n", resp)
}
```

### OpenCorpAPI example usage

```go
package main

func main() {
    cfg, err := wework.NewConfigBuilder().OpenCorp(wework.OpenCorp{
        ProviderCorpID:   "WEWORK_PROVIDER_CORP_ID",
        ProviderSecret:   "WEWORK_PROVIDER_SECRET",
        SuiteID:          "WEWORK_SUITE_ID",
        SuiteSecret:      "WEWORK_SUITE_SECRET",
        SuiteToken:       "WEWORK_SUITE_TOKEN",
        SuiteEncodingAES: "WEWORK_SUITE_AES_KEY",
        AuthCorpID:       "WEWORK_AUTH_CORP_ID",
    }).Build()

    if err != nil {
        fmt.Printf("OpenCorp config build error: %v\n", err)
        return
    }

    client := wework.NewClient(cfg)

    resp, err := client.GetAuthInfo()
    if err != nil {
        fmt.Printf("GetAuthInfo error: %v\n", err)
        return
    }

    fmt.Printf("GetAuthInfo result: %v\n", resp)
}
```

### Use custom storage to store token

```go
package main

type CustomRedisStore struct {
}

func (crs *CustomRedisStore) GetToken(c *wework.Client, tokenType wework.TokenType) (string, error) {
    // redis get token
    return "token...", nil
}

func (crs *CustomRedisStore) SetToken(c *wework.Client, tokenType wework.TokenType, token string, expiresIn int64) (string, error) {
    // redis set token
    return nil
}

func main() {
    cfg, err := wework.NewConfigBuilder().OpenCorp(wework.OpenCorp{
        ProviderCorpID:   "WEWORK_PROVIDER_CORP_ID",
        ProviderSecret:   "WEWORK_PROVIDER_SECRET",
        SuiteID:          "WEWORK_SUITE_ID",
        SuiteSecret:      "WEWORK_SUITE_SECRET",
        SuiteToken:       "WEWORK_SUITE_TOKEN",
        SuiteEncodingAES: "WEWORK_SUITE_AES_KEY",
        AuthCorpID:       "WEWORK_AUTH_CORP_ID",
    }).TokenStore(&CustomRedisStore{}).Build()

    if err != nil {
        fmt.Printf("OpenCorp config build error: %v\n", err)
        return
    }

    client := wework.NewClient(cfg)

    resp, err := client.GetAuthInfo()
    if err != nil {
        fmt.Printf("GetAuthInfo error: %v\n", err)
        return
    }

    fmt.Printf("GetAuthInfo result: %v\n", resp)
}
```
