package basic

import "github.com/jasonwwl/go-wework"

type CorpIDToOpenCorpIDResponse struct {
	wework.APIBaseResponse
	OpenCorpID string `json:"open_corpid"`
}

type UserIDToOpenUserIDResponse struct {
	wework.APIBaseResponse
	OpenUserIDList []struct {
		UserID     string `json:"userid"`
		OpenUserID string `json:"open_userid"`
	} `json:"open_userid_list"`
	InvalidUserIDList []string `json:"invalid_userid_list"`
}

type GetNewExternalUserIDResponse struct {
	wework.APIBaseResponse
	Items []struct {
		ExternalUserID    string `json:"external_userid"`
		NewExternalUserID string `json:"new_external_userid"`
	} `json:"items"`
}

type GetNewGroupChatExternalUserIDResponse struct {
	wework.APIBaseResponse
	Items []struct {
		ExternalUserID    string `json:"external_userid"`
		NewExternalUserID string `json:"new_external_userid"`
	} `json:"items"`
}

type UnionIDToExternalUserIDRequest struct {
	UnionID     string `json:"unionid"`
	OpenID      string `json:"openid"`
	SubjectType int    `json:"subject_type"` //小程序或公众号的主体类型：0表示主体名称是企业的 (默认)，1表示主体名称是服务商的
}

type UnionIDToExternalUserIDResponse struct {
	wework.APIBaseResponse
	ExternalUserID string `json:"external_userid"`
	PendingID      string `json:"pending_id"`
}

type ExternalUserIDToPendingIDRequest struct {
	ChatID         string   `json:"chat_id"`
	ExternalUserID []string `json:"external_userid"`
}

type ExternalUserIDToPendingIDResponse struct {
	wework.APIBaseResponse
	Result []struct {
		ExternalUserID string `json:"external_userid"`
		PendingID      string `json:"pending_id"`
	} `json:"result"`
}

type OpenUserIDToUserIDResponse struct {
	wework.APIBaseResponse
	UserIDList []struct {
		OpenUserID string `json:"open_userid"`
		UserID     string `json:"userid"`
	} `json:"userid_list"`
	InvalidOpenUserIDList []string `json:"invalid_open_userid_list"`
}

type FromServiceExternalUserIDResponse struct {
	wework.APIBaseResponse
	ExternalUserID string `json:"external_userid"`
}
