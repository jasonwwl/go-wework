package basic

import (
	"context"

	"github.com/jasonwwl/go-wework"
)

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
func (c *BasicClient) GetSuiteToken(ctx context.Context, ticket string) (response SuiteTokenResponse, err error) {
	err = c.client.Request(ctx, "POST", "/service/get_suite_token", &response, wework.WithJSONData(&SuiteTokenRequest{
		SuiteID:     c.client.GetOpenCorpConfig().SuiteID,
		SuiteSecret: c.client.GetOpenCorpConfig().SuiteSecret,
		SuiteTicket: ticket,
	}))
	return
}

// 获取预授权码
//
// 该API用于获取预授权码。预授权码用于企业授权时的第三方服务商安全验证。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90601
func (c *BasicClient) GetPreAuthCode(ctx context.Context) (response PreAuthCodeResponse, err error) {
	err = c.client.Request(ctx, "GET", "/service/get_pre_auth_code", &response, wework.WithToken(wework.SuiteToken))
	return
}

// 设置授权配置，该接口可对某次授权进行配置。可支持测试模式（应用未发布时）。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90602
func (c *BasicClient) SetSessionInfo(ctx context.Context, preAuthCode string, sessinInfo SetSessionInfoRequest) (response wework.APIBaseResponse, err error) {

	err = c.client.Request(
		ctx,
		"POST",
		"/service/set_session_info",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(wework.H{
			"pre_auth_code": preAuthCode,
			"session_info":  sessinInfo,
		}),
	)
	return
}

// 获取企业永久授权码
//
// 该API用于使用临时授权码换取授权方的永久授权码，并换取授权信息、企业access_token，临时授权码一次有效。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90603
func (c *BasicClient) GetPermanentCode(ctx context.Context, authCode string) (response GetPermanentCodeResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/service/get_permanent_code",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(wework.H{
			"auth_code": authCode,
		}),
	)
	return
}

// 获取企业授权信息
//
// 该API用于通过永久授权码换取企业微信的授权信息。
// 永久code的获取，是通过临时授权码使用get_permanent_code接口获取到的permanent_code。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90604
func (c *BasicClient) GetAuthInfo(ctx context.Context, authCorpid string, permanentCode string) (response GetPermanentCodeResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/service/get_auth_info",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(wework.H{
			"auth_corpid":    authCorpid,
			"permanent_code": permanentCode,
		}),
	)
	return
}

// 获取企业凭证
//
// 第三方服务商在取得企业的永久授权码后，通过此接口可以获取到企业的access_token。
// 获取后可通过通讯录、应用、消息等企业接口来运营这些应用。
//
// > 此处获得的企业access_token与企业获取access_token拿到的token，本质上是一样的，只不过获取方式不同。获取之后，就跟普通企业一样使用token调用API接口
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90605
func (c *BasicClient) GetCorpToken(ctx context.Context, authCorpid string, permanentCode string) (response GetCorpTokenResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/service/get_corp_token",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(wework.H{
			"auth_corpid":    authCorpid,
			"permanent_code": permanentCode,
		}),
	)
	return
}

// 获取应用的管理员列表
//
// 第三方服务商可以用此接口获取授权企业中某个第三方应用的管理员列表(不包括外部管理员)，
// 以便服务商在用户进入应用主页之后根据是否管理员身份做权限的区分。
//
// > 该应用必须与SUITE_ACCESS_TOKEN对应的suiteid对应，否则没权限查看
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/90606
func (c *BasicClient) GetAdminList(ctx context.Context, authCorpid string, agentid int) (response GetAdminListResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/service/get_admin_list",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(wework.H{
			"auth_corpid": authCorpid,
			"agentid":     agentid,
		}),
	)
	return
}

// GetAppQRCode 获取应用二维码
//
// 用于获取第三方应用二维码。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95430
func (c *BasicClient) GetAppQRCode(ctx context.Context, request GetAppQRCodeRequest) (response GetAppQRCodeResponse, err error) {
	if request.SuiteID == "" {
		request.SuiteID = c.client.GetOpenCorpConfig().SuiteID
	}
	payload := struct {
		ResultType int `json:"result_type"`
		GetAppQRCodeRequest
	}{
		ResultType:          2,
		GetAppQRCodeRequest: request,
	}

	err = c.client.Request(
		ctx,
		"POST",
		"/service/get_app_qrcode",
		&response,
		wework.WithToken(wework.SuiteToken),
		wework.WithJSONData(payload),
	)
	return
}
