package wework

import (
	"context"
	"fmt"
	"net/url"
)

type AccessTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type ProviderTokenRequest struct {
	CorpID         string `json:"corpid"`
	ProviderSecret string `json:"provider_secret"`
}

type ProviderTokenResonspe struct {
	ErrCode             int    `json:"errcode"`
	ErrMsg              string `json:"errmsg"`
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int64  `json:"expires_in"`
}

type SuiteTokenRequest struct {
	SuiteID     string `json:"suite_id"`
	SuiteSecret string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

type SuiteTokenResponse struct {
	ErrCode          int    `json:"errcode"`
	ErrMsg           string `json:"errmsg"`
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int64  `json:"expires_in"`
}

type AuthCorpAccessTokenRequest struct {
	AuthCorpID    string `json:"auth_corpid"`
	PermanentCode string `json:"permanent_code"`
}

type AuthCorpAccessTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetToken 获取指定类型的 Token
func (c *Client) GetToken(ctx context.Context, token *TokenDescriptor) (string, error) {
	tk, err := c.GetStore().GetToken(c, token.TokenType)

	if err != nil && err != ErrNilStoreToken {
		return "", err
	}

	if tk != "" {
		return tk, nil
	}

	switch token.TokenType {
	case AccessToken.TokenType:
		return c.FetchAccessToken(ctx)
	case ProviderToken.TokenType:
		return c.FetchProviderToken(ctx)
	case SuiteToken.TokenType:
		return c.FetchSuiteToken(ctx)
	case AuthCorpAccessToken.TokenType:
		return c.FetchAuthCorpAccessToken(ctx)
	default:
		return "", fmt.Errorf("invalid token type: %s", token.TokenType)
	}
}

// FetchAccessToken 获取企业微信 AccessToken
func (c *Client) FetchAccessToken(ctx context.Context) (string, error) {
	query := url.Values{}
	query.Add("corpid", c.config.InternalCorp.CorpID)
	query.Add("corpsecret", c.config.InternalCorp.Secret)

	req, err := c.NewRequest(ctx, "GET", "/gettoken", WithQuery(query))
	if err != nil {
		return "", err
	}

	var resp AccessTokenResponse
	err = c.SendRequest(req, &resp)
	if err != nil {
		return "", err
	}

	err = c.config.Options.TokenStore.SetToken(c, AccessToken.TokenType, resp.AccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}

// FetchProviderToken 获取服务商 AccessToken
func (c *Client) FetchProviderToken(ctx context.Context) (string, error) {

	req, err := c.NewRequest(ctx, "POST", "/service/get_provider_token", WithJSONData(&ProviderTokenRequest{
		CorpID:         c.config.OpenCorp.ProviderCorpID,
		ProviderSecret: c.config.OpenCorp.ProviderSecret,
	}))

	if err != nil {
		return "", err
	}

	var resp ProviderTokenResonspe
	err = c.SendRequest(req, &resp)
	if err != nil {
		return "", err
	}

	err = c.GetStore().SetToken(c, ProviderToken.TokenType, resp.ProviderAccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	return resp.ProviderAccessToken, nil
}

// FetchSuiteToken 获取第三方应用 AccessToken
func (c *Client) FetchSuiteToken(ctx context.Context) (string, error) {
	ticket, err := c.GetStore().GetToken(c, SuiteTicket.TokenType)

	if err != nil {
		return "", err
	}

	req, err := c.NewRequest(ctx, "POST", "/service/get_suite_token", WithJSONData(&SuiteTokenRequest{
		SuiteID:     c.config.OpenCorp.SuiteID,
		SuiteSecret: c.config.OpenCorp.SuiteSecret,
		SuiteTicket: ticket,
	}))

	if err != nil {
		return "", err
	}

	var resp SuiteTokenResponse
	err = c.SendRequest(req, &resp)
	if err != nil {
		return "", err
	}

	err = c.GetStore().SetToken(c, SuiteToken.TokenType, resp.SuiteAccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	return resp.SuiteAccessToken, nil
}

// FetchAuthCorpAccessToken 获取授权企业 AccessToken
func (c *Client) FetchAuthCorpAccessToken(ctx context.Context) (string, error) {
	store := c.GetStore()

	pCode, err := store.GetToken(c, PermanentCode.TokenType)
	if err != nil {
		return "", err
	}

	req, err := c.NewRequest(ctx, "POST", "/service/get_corp_token", WithJSONData(&AuthCorpAccessTokenRequest{
		AuthCorpID:    c.config.OpenCorp.AuthCorpID,
		PermanentCode: pCode,
	}))

	if err != nil {
		return "", err
	}

	var resp AuthCorpAccessTokenResponse
	err = c.SendRequest(req, &resp)
	if err != nil {
		return "", err
	}

	err = store.SetToken(c, ProviderToken.TokenType, resp.AccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
