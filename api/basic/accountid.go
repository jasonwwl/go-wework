package basic

import (
	"context"

	"github.com/jasonwwl/go-wework"
)

// 用于将企业主体的明文corpid转换为服务商主体的密文corpid。
//
// NOTE: 该接口仅适用于第三方服务商
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/97061
func (c *BasicClient) CropIDToOpenCorpID(ctx context.Context, corpid string) (response CorpIDToOpenCorpIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/service/corpid_to_opencorpid",
		&response,
		wework.WithJSONData(wework.H{
			"corpid": corpid,
		}),
		wework.WithToken(wework.ProviderToken),
	)
	return
}

// 将企业主体下的明文userid转换为服务商主体下的密文userid。
//
// NOTE: 该接口仅适用于第三方服务商
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/97062
func (c *BasicClient) UserIDToOpenUserID(ctx context.Context, useridList []string) (response UserIDToOpenUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/batch/userid_to_openuserid",
		&response,
		wework.WithJSONData(wework.H{
			"userid_list": useridList,
		}),
		wework.WithToken(wework.AuthCorpAccessToken),
	)
	return
}

// 转换客户external_userid
//
// 将企业主体下的external_userid转换为服务商主体下的external_userid。
//
// NOTE: 该接口仅适用于第三方服务商
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/97063
func (c *BasicClient) GetNewExternalUserID(ctx context.Context, externalUseridList []string) (response GetNewExternalUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/externalcontact/get_new_external_userid",
		&response,
		wework.WithJSONData(wework.H{
			"external_userid_list": externalUseridList,
		}),
		wework.WithToken(wework.AuthCorpAccessToken),
	)
	return
}

// 转换客户群成员external_userid
//
// 转换客户external_userid接口(GetNewExternalUserID)不支持客户群的场景，
// 如果需要转换客户群中无好友关系的群成员external_userid，
// 需要调用本接口，调用时需要传入客户群的chat_id。
//
// NOTE: 该接口仅适用于第三方服务商
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/97063
func (c *BasicClient) GetNewGroupChatExternalUserID(ctx context.Context, chatId string, externalUseridList []string) (response GetNewGroupChatExternalUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/externalcontact/groupchat/get_new_external_userid",
		&response,
		wework.WithJSONData(wework.H{
			"chat_id":              chatId,
			"external_userid_list": externalUseridList,
		}),
		wework.WithToken(wework.AuthCorpAccessToken),
	)
	return
}

// unionid转换为第三方external_userid
//
// 当微信用户进入服务商的小程序或公众号时，服务商可通过此接口，将微信客户的unionid转为第三方主体的external_userid，若该微信用户尚未成为企业的客户，则返回pending_id。
// 小程序或公众号的主体名称可以是企业的，也可以是服务商的。

// 该接口有调用频率限制，当subject_type为0时，按企业作如下的限制：10万次/小时、48万次/天、750万次/月；（注意这里是所有服务商共用企业额度的）
// 当subject_type为1时，按服务商作如下的限制：10万次/小时、48万次/天、750万次/月
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95900
func (c *BasicClient) UnionIDToExternalUserID(ctx context.Context, request UnionIDToExternalUserIDRequest) (response UnionIDToExternalUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/idconvert/unionid_to_external_userid",
		&response,
		wework.WithJSONData(request),
		wework.WithToken(wework.AccessToken),
	)
	return
}

// external_userid查询pending_id
//
// 该接口可用于当一个微信用户成为企业客户前已经使用过服务商服务（服务商曾通过unionid查询external_userid接口获取到pending_id）的场景。
// 本接口获取到的pending_id可以维持unionid和external_userid的关联关系。
// pending_id有效期为90天，超过有效期之后，将无法通过该接口将external_userid换取对应的pending_id。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95900
func (c *BasicClient) ExternalUserIDToPendingID(ctx context.Context, request ExternalUserIDToPendingIDRequest) (response ExternalUserIDToPendingIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/idconvert/batch/external_userid_to_pending_id",
		&response,
		wework.WithJSONData(request),
		wework.WithToken(wework.AccessToken),
	)
	return
}

// 将代开发应用或第三方应用获取的密文open_userid转换为明文userid。
//
// NOTE: 该接口仅适用于企业内部开发
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95884
func (c *BasicClient) OpenUserIDToUserID(ctx context.Context, openUserIDList []string) (response OpenUserIDToUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/cgi-bin/batch/openuserid_to_userid",
		&response,
		wework.WithJSONData(wework.H{
			"open_userid_list": openUserIDList,
			"source_agentid":   c.client.GetInternalCorpConfig().AgentID,
		}),
		wework.WithToken(wework.AccessToken),
	)
	return
}

// 将代开发应用或第三方应用获取的externaluserid转换成自建应用的externaluserid。
//
// NOTE: 该接口仅适用于企业内部开发
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95884
func (c *BasicClient) FromServiceExternalUserID(ctx context.Context, externalUserID string) (response FromServiceExternalUserIDResponse, err error) {
	err = c.client.Request(
		ctx,
		"POST",
		"/externalcontact/from_service_external_userid",
		&response,
		wework.WithJSONData(wework.H{
			"external_userid": externalUserID,
			"source_agentid":  c.client.GetInternalCorpConfig().AgentID,
		}),
		wework.WithToken(wework.AccessToken),
	)
	return
}
