package wework

import (
	"net/http"
	"time"
)

type configBuilder struct {
	config *ClientConfig
}

func NewConfigBuilder() *configBuilder {
	return &configBuilder{&ClientConfig{
		Options: &Options{
			Timeout: Timeout,
			BaseURL: WeWorkAPIURL,
		},
		HTTPClient: &http.Client{},
	}}
}

func (b *configBuilder) Timeout(timeout time.Duration) *configBuilder {
	b.config.Options.Timeout = timeout
	b.config.HTTPClient.Timeout = timeout
	return b
}

func (b *configBuilder) TokenStore(store Store) *configBuilder {
	b.config.Options.TokenStore = store
	return b
}

func (b *configBuilder) BaseURL(baseURL string) *configBuilder {
	b.config.Options.BaseURL = baseURL
	return b
}

func (b *configBuilder) HTTPClient(client *http.Client) *configBuilder {
	b.config.HTTPClient = client
	return b
}

func (b *configBuilder) OpenCorp(openCorpCfg OpenCorp) *configBuilder {
	b.config.OpenCorp = &openCorpCfg
	return b
}

func (b *configBuilder) InternalCorp(internalCorpCfg InternalCorp) *configBuilder {
	b.config.InternalCorp = &internalCorpCfg
	return b
}

func (b *configBuilder) DebugMode(debugMode bool) *configBuilder {
	b.config.DebugMode = debugMode
	return b
}

func (b *configBuilder) Build() (*ClientConfig, error) {
	if b.config.Options.TokenStore == nil {
		b.config.Options.TokenStore = InitMemoryStore()
	}

	return b.config, nil
}
