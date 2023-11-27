package wework

import (
	"context"
	"fmt"
	"net/url"
)

// GetToken 获取指定类型的 Token
func (c *Client) GetToken(ctx context.Context, token *TokenDescriptor) (tk string, err error) {

	internalCorpCfg := c.GetInternalCorpConfig()
	openCorpCfg := c.GetOpenCorpConfig()

	if internalCorpCfg == nil && openCorpCfg == nil {
		return "", fmt.Errorf("invalid config: internal corp config and open corp config are both nil")
	}

	if internalCorpCfg != nil && openCorpCfg != nil {
		return "", fmt.Errorf("invalid config: internal corp config and open corp config are both not nil")
	}

	tokenType := token.TokenType

	if tokenType == AccessToken.TokenType || tokenType == AuthCorpAccessToken.TokenType {
		if internalCorpCfg != nil {
			tokenType = AccessToken.TokenType
		}

		if openCorpCfg != nil {
			tokenType = AuthCorpAccessToken.TokenType
		}
	}

	switch tokenType {
	case AccessToken.TokenType:
		return c.FetchAccessTokenIfNeeded(ctx)
	case AuthCorpAccessToken.TokenType:
		return c.FetchAuthCorpAccessTokenIfNeeded(ctx)
	case ProviderToken.TokenType:
		return c.FetchProviderTokenIfNeeded(ctx)
	case SuiteToken.TokenType:
		return c.FetchSuiteTokenIfNeeded(ctx)
	default:
		return "", fmt.Errorf("invalid token type: %s", token.TokenType)
	}
}

// 从缓存中获取AccessToken（内部应用），如果缓存中没有，则从企微API获取
func (c *Client) FetchAccessTokenIfNeeded(ctx context.Context) (tk string, err error) {
	internalCorpCfg := c.GetInternalCorpConfig()
	if internalCorpCfg == nil {
		err = fmt.Errorf("invalid config: internal corp config is nil")
		return
	}

	tk, err = c.GetStore().GetToken(c, ctx, AccessToken.TokenType)
	if err == nil && tk != "" {
		return
	}

	resp, err := c.GetAccessToken(ctx, internalCorpCfg.CorpID, internalCorpCfg.Secret)
	if err != nil {
		return
	}

	tk = resp.AccessToken
	err = c.GetStore().SetToken(c, ctx, AccessToken.TokenType, tk, resp.ExpiresIn)

	return
}

// 从缓存中获取AccessToken（第三方授权应用），如果缓存中没有，则从企微API获取
func (c *Client) FetchAuthCorpAccessTokenIfNeeded(ctx context.Context) (tk string, err error) {
	openCorpCfg := c.GetOpenCorpConfig()
	if openCorpCfg == nil {
		err = fmt.Errorf("invalid config: open corp config is nil")
		return
	}

	tk, err = c.GetStore().GetToken(c, ctx, AuthCorpAccessToken.TokenType)
	if err == nil && tk != "" {
		return
	}

	permanentCode, err := c.GetStore().GetToken(c, ctx, PermanentCode.TokenType)
	if err != nil {
		return
	}

	resp, err := c.GetAuthCorpAccessToken(ctx, openCorpCfg.AuthCorpID, permanentCode)
	if err != nil {
		return
	}

	tk = resp.AccessToken
	err = c.GetStore().SetToken(c, ctx, AuthCorpAccessToken.TokenType, tk, resp.ExpiresIn)

	return
}

// 获取服务商凭证，如果缓存中没有，则从企微API获取
func (c *Client) FetchProviderTokenIfNeeded(ctx context.Context) (tk string, err error) {

	openCorpCfg := c.GetOpenCorpConfig()
	if openCorpCfg == nil {
		err = fmt.Errorf("invalid config: open corp config is nil")
		return
	}

	tk, err = c.GetStore().GetToken(c, ctx, ProviderToken.TokenType)
	if err == nil && tk != "" {
		return
	}

	resp, err := c.GetProviderToken(ctx, openCorpCfg.ProviderCorpID, openCorpCfg.ProviderSecret)
	if err != nil {
		return
	}

	err = c.GetStore().SetToken(c, ctx, ProviderToken.TokenType, resp.ProviderAccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	tk = resp.ProviderAccessToken
	return
}

// 获取第三方应用凭证，如果缓存中没有，则从企微API获取
func (c *Client) FetchSuiteTokenIfNeeded(ctx context.Context) (tk string, err error) {
	openCorpCfg := c.GetOpenCorpConfig()
	if openCorpCfg == nil {
		err = fmt.Errorf("invalid config: open corp config is nil")
		return
	}

	tk, err = c.GetStore().GetToken(c, ctx, SuiteTicket.TokenType)
	if err == nil && tk != "" {
		return
	}

	ticket, err := c.GetStore().GetToken(c, ctx, SuiteTicket.TokenType)
	if err != nil {
		return
	}

	resp, err := c.GetSuiteToken(ctx, openCorpCfg.SuiteID, openCorpCfg.SuiteSecret, ticket)
	if err != nil {
		return
	}

	err = c.GetStore().SetToken(c, ctx, SuiteTicket.TokenType, resp.SuiteAccessToken, resp.ExpiresIn)
	if err != nil {
		return "", err
	}

	tk = resp.SuiteAccessToken
	return
}

// 获取服务商凭证
//
// 该API用于获取服务商凭证，该凭证用于服务商调用企业微信开放接口。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/91200
func (c *Client) GetProviderToken(ctx context.Context, corpid string, providerSecret string) (response GetProviderTokenResponse, err error) {
	err = c.Request(
		ctx,
		"POST",
		"/service/get_provider_token",
		&response,
		WithJSONData(H{
			"corpid":          corpid,
			"provider_secret": providerSecret,
		}),
	)
	return
}

// 内部应用获取access_token
//
// 获取access_token是调用企业微信API接口的第一步，相当于创建了一个登录凭证，其它的业务API接口，都需要依赖于access_token来鉴权调用者身份。
// 因此开发者，在使用业务接口前，要明确access_token的颁发来源，使用正确的access_token。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/91039
func (c *Client) GetAccessToken(ctx context.Context, corpid string, corpSecret string) (response GetTokenResponse, err error) {
	query := url.Values{}
	query.Add("corpid", corpid)
	query.Add("corpsecret", corpSecret)
	err = c.Request(
		ctx,
		"GET",
		"/gettoken",
		&response,
		WithQuery(query),
	)
	return
}

// 第三方服务商获取授权企业的access_token
//
// 第三方服务商在取得企业的永久授权码后，通过此接口可以获取到企业的access_token。
// 获取后可通过通讯录、应用、消息等企业接口来运营这些应用。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90605
//   - 此处获得的企业access_token与[企业获取access_token]拿到的token，本质上是一样的，只不过获取方式不同。获取之后，就跟普通企业一样使用token调用API接口
//
// [企业获取access_token]: https://developer.work.weixin.qq.com/document/10013
func (c *Client) GetAuthCorpAccessToken(ctx context.Context, authCorpid, permanentCode string) (response AuthCorpAccessTokenResponse, err error) {
	err = c.Request(
		ctx,
		"POST",
		"/service/get_corp_token",
		&response,
		WithJSONData(H{
			"auth_corpid":    authCorpid,
			"permanent_code": permanentCode,
		}),
		WithToken(SuiteToken),
	)
	return
}

// 获取第三方应用凭证
//
// 该API用于获取第三方应用凭证（suite_access_token）。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90600
//   - 由于第三方服务商可能托管了大量的企业，其安全问题造成的影响会更加严重，故API中除了合法来源IP校验之外，还额外增加了suite_ticket作为安全凭证。
//   - 获取suite_access_token时，需要suite_ticket参数。suite_ticket由企业微信后台定时推送给“指令回调URL”，每十分钟更新一次，见[推送suite_ticket]。
//   - suite_ticket实际有效期为30分钟，可以容错连续两次获取suite_ticket失败的情况，但是请永远使用最新接收到的suite_ticket。通过本接口获取的suite_access_token有效期为2小时，开发者需要进行缓存，不可频繁获取。
//
// [推送suite_ticket]: https://developer.work.weixin.qq.com/document/path/90628
func (c *Client) GetSuiteToken(ctx context.Context, suiteId, suiteSecret, suiteTicket string) (response SuiteTokenResponse, err error) {
	err = c.Request(
		ctx,
		"POST",
		"/service/get_suite_token",
		&response,
		WithJSONData(H{
			"suite_id":     suiteId,
			"suite_secret": suiteSecret,
			"suite_ticket": suiteTicket,
		}),
	)
	return
}
