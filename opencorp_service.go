package wework

import (
	"context"
)

type PreAuthCodeResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

// GetPreAuthCode 获取预授权码
//
// 该API用于获取预授权码。预授权码用于企业授权时的第三方服务商安全验证。
// https://developer.work.weixin.qq.com/document/path/90601
func (c *Client) GetPreAuthCode(ctx context.Context) (response PreAuthCodeResponse, err error) {
	req, err := c.NewRequest(ctx, "GET", "/service/get_pre_auth_code", WithToken(SuiteToken))
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}

type SessionInfoRequest struct {
	PreAuthCode string `json:"pre_auth_code"`
	SessionInfo struct {
		AppID    string `json:"appid"`
		AuthType int    `json:"auth_type"`
	} `json:"session_info"`
}

type SessionInfoResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetSessionInfo 设置授权配置
//
// 该接口可对某次授权进行配置。可支持测试模式（应用未发布时）。
// https://developer.work.weixin.qq.com/document/path/90602
func (c *Client) GetSessionInfo(ctx context.Context, request SessionInfoRequest) (response SessionInfoResponse, err error) {
	req, err := c.NewRequest(
		ctx,
		"POST",
		"/service/get_pre_auth_code",
		WithToken(SuiteToken),
		WithJSONData(request),
	)
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}

type GetPermanentCodeResponse struct {
	ErrCode          int              `json:"errcode"`
	ErrMsg           string           `json:"errmsg"`
	AccessToken      string           `json:"access_token"`
	ExpiresIn        int64            `json:"expires_in"`
	PermanentCode    string           `json:"permanent_code"`
	DealerCorpInfo   DealerCorpInfo   `json:"dealer_corp_info"`
	AuthCorpInfo     AuthCorpInfo     `json:"auth_corp_info"`
	AuthInfo         AuthInfo         `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
	State            string           `json:"state"`
}

type DealerCorpInfo struct {
	CorpID   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

type AuthCorpInfo struct {
	CorpID            string `json:"corpid"`
	CorpName          string `json:"corp_name"`
	CorpType          string `json:"corp_type"`
	CorpSquareLogoURL string `json:"corp_square_logo_url"`
	CorpUserMax       int    `json:"corp_user_max"`
	CorpFullName      string `json:"corp_full_name"`
	VerifiedEndTime   int64  `json:"verified_end_time"`
	SubjectType       int    `json:"subject_type"`
	CorpWXQrcode      string `json:"corp_wxqrcode"`
	CorpScale         string `json:"corp_scale"`
	CorpIndustry      string `json:"corp_industry"`
	CorpSubIndustry   string `json:"corp_sub_industry"`
}

type AuthInfo struct {
	Agent []AuthInfoAgent `json:"agent"`
}

type AuthInfoAgent struct {
	AgentID          int             `json:"agentid"`
	Name             string          `json:"name"`
	RoundLogoURL     string          `json:"round_logo_url"`
	SquareLogoURL    string          `json:"square_logo_url"`
	AppID            int             `json:"appid"`
	AuthMode         int             `json:"auth_mode"`
	IsCustomizedApp  bool            `json:"is_customized_app"`
	AuthFromThirdApp bool            `json:"auth_from_thirdapp"`
	Privilege        AgentPrivilege  `json:"privilege"`
	SharedFrom       AgentSharedFrom `json:"shared_from"`
}

type AgentPrivilege struct {
	Level      int      `json:"level"`
	AllowParty []int    `json:"allow_party"`
	AllowUser  []string `json:"allow_user"`
	AllowTag   []int    `json:"allow_tag"`
	ExtraParty []int    `json:"extra_party"`
	ExtraUser  []string `json:"extra_user"`
	ExtraTag   []int    `json:"extra_tag"`
}

type AgentSharedFrom struct {
	CorpID    string `json:"corpid"`
	ShareType int    `json:"share_type"`
}

type AuthUserInfo struct {
	UserID     string `json:"userid"`
	OpenUserID string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateID   string `json:"template_id"`
	State        string `json:"state"`
}

// GetPermanentCode 获取企业永久授权码
//
// 该API用于使用临时授权码换取授权方的永久授权码，并换取授权信息、企业access_token，临时授权码一次有效。
// https://developer.work.weixin.qq.com/document/path/90603
func (c *Client) GetPermanentCode(ctx context.Context, authCode string) (response GetPermanentCodeResponse, err error) {
	req, err := c.NewRequest(
		ctx,
		"POST",
		"/service/get_permanent_code",
		WithToken(SuiteToken),
		WithJSONData(map[string]string{
			"auth_code": authCode,
		}),
	)
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}

type GetAuthInfoResponse struct {
	ErrCode        int            `json:"errcode"`
	ErrMsg         string         `json:"errmsg"`
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       AuthInfo       `json:"auth_info"`
}

// GetAuthInfo 获取企业授权信息
//
// 该API用于通过永久授权码换取企业微信的授权信息。 永久code的获取，是通过临时授权码使用get_permanent_code 接口获取到的permanent_code。
// https://developer.work.weixin.qq.com/document/path/90604
func (c *Client) GetAuthInfo(ctx context.Context) (response GetAuthInfoResponse, err error) {
	pCode, err := c.GetToken(ctx, PermanentCode)
	if err != nil {
		return
	}

	req, err := c.NewRequest(
		ctx,
		"POST",
		"/service/get_auth_info",
		WithToken(SuiteToken),
		WithJSONData(map[string]string{
			"auth_corpid":    c.GetOpenCorpConfig().AuthCorpID,
			"permanent_code": pCode,
		}),
	)
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}

type GetAdminListResponse struct {
	ErrCode   int         `json:"errcode"`
	ErrMsg    string      `json:"errmsg"`
	AdminList []AdminList `json:"admin_list"`
}

type AdminList struct {
	UserID     string `json:"userid"`
	OpenUserID string `json:"open_userid"`
	AuthType   int    `json:"auth_type"`
}

// GetAdminList 获取应用的管理员列表
//
// 第三方服务商可以用此接口获取授权企业中某个第三方应用的管理员列表(不包括外部管理员)，以便服务商在用户进入应用主页之后根据是否管理员身份做权限的区分。
//
// > 该应用必须与SUITE_ACCESS_TOKEN对应的suiteid对应，否则没权限查看
//
// https://work.weixin.qq.com/api/doc/90001/90143/91120
func (c *Client) GetAdminList(ctx context.Context, agentId int) (response GetAdminListResponse, err error) {
	req, err := c.NewRequest(
		ctx,
		"POST",
		"/service/get_auth_info",
		WithToken(SuiteToken),
		WithJSONData(map[string]any{
			"auth_corpid": c.GetOpenCorpConfig().AuthCorpID,
			"agentid":     agentId,
		}),
	)
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}

type GetAppQRCodeRequest struct {
	AppID int    `json:"appid,omitempty"`
	State string `json:"state,omitempty"`
	Style int    `json:"style,omitempty"`
}

type GetAppQRCodeResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	QRCode  string `json:"qrcode"`
}

// GetAppQRCode 获取应用二维码
//
// 用于获取第三方应用二维码。
// https://developer.work.weixin.qq.com/document/path/95430
func (c *Client) GetAppQRCode(ctx context.Context, request GetAppQRCodeRequest) (response GetAppQRCodeResponse, err error) {
	payload := struct {
		SuiteId    string `json:"suite_id"`
		ResultType int    `json:"result_type"`
		GetAppQRCodeRequest
	}{
		SuiteId:             c.GetOpenCorpConfig().SuiteID,
		ResultType:          2,
		GetAppQRCodeRequest: request,
	}

	req, err := c.NewRequest(
		ctx,
		"POST",
		"/service/get_app_qrcode",
		WithToken(SuiteToken),
		WithJSONData(payload),
	)
	if err != nil {
		return
	}

	err = c.SendRequest(req, &response)
	if err != nil {
		return
	}

	return
}
