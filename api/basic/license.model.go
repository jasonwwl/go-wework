package basic

import "github.com/jasonwwl/go-wework"

type ListOrderAccountRequest struct {
	OrderID string `json:"corpid"` // 订单id
	Limit   int    `json:"limit"`  // 返回的最大记录数，整型，最大值1000，默认值500
	Cursor  string `json:"cursor"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
}

type ListOrderAccountResponse struct {
	wework.APIBaseResponse
	NextCursor  string `json:"next_cursor"` // 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回""
	HasMore     int    `json:"has_more"`    // 是否有更多。 0: 没有， 1: 有
	AccountList []struct {
		ActiveCode string `json:"active_code"` // 账号码，订单类型为购买账号时，返回该字段
		UserID     string `json:"userid"`      // 企业续期成员userid，订单类型为续期账号时，返回该字段。返回加密的userid
		Type       int    `json:"type"`        // 账号类型，1-基础账号；2-互通账号
	} `json:"account_list"` // 订单中的账号列表
}

// 需要激活的账号信息
type ActiveListItem struct {
	ActiveCode string `json:"active_code"` // 账号激活码
	UserID     string `json:"userid"`      // 待绑定激活的企业成员userid
}

// 批量激活账号结果
type BatchActiveAccountResponse struct {
	wework.APIBaseResponse
	ActiveList []struct {
		ActiveCode string `json:"active_code"` // 账号激活码
		UserID     string `json:"userid"`      // 待绑定激活的企业成员userid
		ErrCode    int    `json:"errcode"`     // 用户激活错误码，0为成功
	} `json:"active_list"`
}

type ListActivedAccountRequest struct {
	CorpID string `json:"corpid"`           // 企业corpid
	Limit  int    `json:"limit,omitempty"`  // 返回的最大记录数，整型，最大值1000，默认值500
	Cursor string `json:"cursor,omitempty"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
}

type ListActivedAccountResponse struct {
	wework.APIBaseResponse
	NextCursor  string `json:"next_cursor"` // 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回""
	HasMore     int    `json:"has_more"`    // 是否还有更多数据, 0-否；1-是
	AccountList []struct {
		UserId     string `json:"userid"`      // 企业成员userid
		Type       int    `json:"type"`        // 账号类型，1-基础账号；2-互通账号
		ExpireTime int64  `json:"expire_time"` // 账号过期时间，单位为秒
		ActiveTime int64  `json:"active_time"` // 账号激活时间，单位为秒
	} `json:"account_list"` // 已激活成员列表，已激活过期的也会返回
}
