package basic

import (
	"context"

	"github.com/jasonwwl/go-wework"
)

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
func (c *BasicClient) GetAuthInfo(ctx context.Context, authCorpid string, permanentCode string) (response GetAuthInfoResponse, err error) {
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

// 获取应用的管理员列表
//
// 第三方服务商可以用此接口获取授权企业中某个第三方应用的管理员列表(不包括外部管理员)，
// 以便服务商在用户进入应用主页之后根据是否管理员身份做权限的区分。
//
//   - 该应用必须与SUITE_ACCESS_TOKEN对应的suiteid对应，否则没权限查看
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
