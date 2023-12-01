package basic

import "github.com/jasonwwl/go-wework"

type SuiteTokenRequest struct {
	SuiteID     string `json:"suite_id"`
	SuiteSecret string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

type SuiteTokenResponse struct {
	wework.APIBaseResponse
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int64  `json:"expires_in"`
}

type PreAuthCodeResponse struct {
	wework.APIBaseResponse
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

type SetSessionInfoRequest struct {
	AppID    []int `json:"appid"`
	AuthType int   `json:"auth_type"`
}

type GetPermanentCodeResponse struct {
	wework.APIBaseResponse
	AccessToken      string           `json:"access_token"`
	ExpiresIn        int64            `json:"expires_in"`
	PermanentCode    string           `json:"permanent_code"`
	DealerCorpInfo   DealerCorpInfo   `json:"dealer_corp_info"`
	AuthCorpInfo     AuthCorpInfo     `json:"auth_corp_info"`
	AuthInfo         AuthInfo         `json:"auth_info"`
	AuthUserInfo     AuthUserInfo     `json:"auth_user_info"`
	RegisterCodeInfo RegisterCodeInfo `json:"register_code_info"`
	State            string           `json:"state"`
	EditionInfo      *EditionInfo     `json:"edition_info"`
}

type GetAuthInfoResponse struct {
	wework.APIBaseResponse
	DealerCorpInfo DealerCorpInfo `json:"dealer_corp_info"`
	AuthCorpInfo   AuthCorpInfo   `json:"auth_corp_info"`
	AuthInfo       AuthInfo       `json:"auth_info"`
	EditionInfo    *EditionInfo   `json:"edition_info"`
}

type GetCorpTokenResponse struct {
	wework.APIBaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type GetAdminListResponse struct {
	wework.APIBaseResponse
	Admin []struct {
		UserID     string `json:"userid"`
		OpenUserID string `json:"open_userid"`
		AuthType   int    `json:"auth_type"`
	} `json:"admin"`
}

type GetAppQRCodeRequest struct {
	SuiteID string `json:"suite_id"`
	AppID   int    `json:"appid,omitempty"`
	State   string `json:"state,omitempty"`
	Style   int    `json:"style,omitempty"`
}

type GetAppQRCodeResponse struct {
	wework.APIBaseResponse
	QRCode string `json:"qrcode"`
}

// 代理服务商企业信息。应用被代理后才有该信息
type DealerCorpInfo struct {
	CorpID   string `json:"corpid"`
	CorpName string `json:"corp_name"`
}

// 授权方企业信息
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

// 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
type AuthInfo struct {
	Agent []AuthInfoAgent `json:"agent"`
}

// 授权的应用信息
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

// 应用对应的权限
type AgentPrivilege struct {
	Level      int      `json:"level"`
	AllowParty []int    `json:"allow_party"`
	AllowUser  []string `json:"allow_user"`
	AllowTag   []int    `json:"allow_tag"`
	ExtraParty []int    `json:"extra_party"`
	ExtraUser  []string `json:"extra_user"`
	ExtraTag   []int    `json:"extra_tag"`
}

// 共享了应用的企业信息，仅当由企业互联或者上下游共享应用触发的安装时才返回
type AgentSharedFrom struct {
	CorpID    string `json:"corpid"`
	ShareType int    `json:"share_type"`
}

// 授权管理员的信息，可能不返回
type AuthUserInfo struct {
	UserID     string `json:"userid"`
	OpenUserID string `json:"open_userid"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
}

// 推广二维码安装相关信息，扫推广二维码安装时返回。成员授权时暂不支持。（注：无论企业是否新注册，只要通过扫推广二维码安装，都会返回该字段）
type RegisterCodeInfo struct {
	RegisterCode string `json:"register_code"`
	TemplateID   string `json:"template_id"`
	State        string `json:"state"`
}

type EditionInfo struct {
	Agent []EditionInfoItem `json:"agent"`
}

type EditionInfoItem struct {
	AgentID               int    `json:"agentid"`
	EditionID             string `json:"edition_id"`
	EditionName           string `json:"edition_name"`
	AppStatus             int    `json:"app_status"`
	UserLimit             int    `json:"user_limit"`
	ExpiredTime           int64  `json:"expired_time"`
	IsVirtualVersion      bool   `json:"is_virtual_version"`
	IsSharedFromOtherCorp bool   `json:"is_shared_from_other_corp"`
}
