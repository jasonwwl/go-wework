package basic

import (
	"context"

	"github.com/jasonwwl/go-wework"
)

func (c *BasicClient) ListOrderAccount(ctx context.Context, request ListOrderAccountRequest) (response ListOrderAccountResponse, err error) {

	err = c.client.Request(
		ctx,
		"POST",
		"/license/list_order_account",
		&response,
		wework.WithToken(wework.ProviderToken),
		wework.WithJSONData(request),
	)
	return
}

// 激活单个账号
//
// 下单购买账号并支付完成之后，先调用获取订单中的账号列表接口获取到账号激活码，
// 然后可以调用该接口将激活码绑定到某个企业员工，以对其激活相应的平台服务能力。
//  1. 一个userid允许激活一个基础账号以及一个互通账号。
//  2. 若userid已激活，使用同类型的激活码来激活后，则绑定关系变为新激活码，新激活码有效时长自动叠加上旧激活码剩余时长，同时旧激活码失效。
//  3. 为了避免服务商调用出错，多个同类型的激活码累加后的有效期不可超过5年，否则接口报错701030。
//  4. 为了避免服务商调用出错，只有当旧的激活码的剩余使用小于等于20天，才可以使用新的同类型的激活码进行激活并续期。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95553
func (c *BasicClient) ActiveAccount(ctx context.Context, code string, corpid string, userid string) (response wework.APIBaseResponse, err error) {

	err = c.client.Request(
		ctx,
		"POST",
		"/license/active_account",
		&response,
		wework.WithToken(wework.ProviderToken),
		wework.WithJSONData(wework.H{
			"active_code": code,
			"corpid":      corpid,
			"userid":      userid,
		}),
	)
	return
}

// 批量激活账号
//
// 可在一次请求里为一个企业的多个成员激活许可账号，便于服务商批量化处理。
//  1. 一个userid允许激活一个基础账号以及一个互通账号。
//  2. 若userid已激活，使用同类型的激活码来激活后，则绑定关系变为新激活码，新激活码有效时长自动叠加上旧激活码剩余时长，同时旧激活码失效。
//  3. 为了避免服务商调用出错，多个同类型的激活码累加后的有效期不可超过5年，否则接口报错701030。
//  4. 为了避免服务商调用出错，只有当旧的激活码的剩余使用小于等于20天，才可以使用新的同类型的激活码进行激活并续期。
//  5. 单次激活的员工数量不超过1000。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95553
func (c *BasicClient) BatchActiveAccount(ctx context.Context, corpid string, activeList []ActiveListItem) (response BatchActiveAccountResponse, err error) {

	err = c.client.Request(
		ctx,
		"POST",
		"/license/batch_active_account",
		&response,
		wework.WithToken(wework.ProviderToken),
		wework.WithJSONData(wework.H{
			"corpid":      corpid,
			"active_list": activeList,
		}),
	)
	return
}

// 获取企业的账号列表
//
// 查询指定企业下的平台能力服务账号列表。
//   - 若为上下游场景，corpid指定的为上游企业，仅返回上游企业激活的账号；若corpid指定为下游企业，若激活码为上游企业分享过来的且已绑定，也会返回。
//
// 文档地址: https://developer.work.weixin.qq.com/document/path/95544
func (c *BasicClient) ListActivedAccount(ctx context.Context, request ListActivedAccountRequest) (response ListActivedAccountResponse, err error) {

	err = c.client.Request(
		ctx,
		"POST",
		"/license/list_actived_account",
		&response,
		wework.WithToken(wework.ProviderToken),
		wework.WithJSONData(request),
	)
	return
}
