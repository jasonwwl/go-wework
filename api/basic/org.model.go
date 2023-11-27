package basic

import "github.com/jasonwwl/go-wework"

type UserInfo struct {
	UserID           string   `json:"userid"`            // 企业微信用户ID
	Name             string   `json:"name"`              // 成员名称
	Department       []int    `json:"department"`        // 成员所属部门id列表
	Order            []int    `json:"order"`             // 部门内的排序值，默认为0，成员次序以创建时间从小到大排列。个数必须和department一致，数值越小排序越靠前。
	Position         string   `json:"position"`          // 职务信息；第三方仅通讯录应用可获取
	Mobile           string   `json:"mobile"`            // 手机号码，第三方仅通讯录应用可获取
	Gender           string   `json:"gender"`            // 性别。0表示未定义，1表示男性，2表示女性
	Email            string   `json:"email"`             // 邮箱，第三方仅通讯录应用可获取
	BizMail          string   `json:"biz_mail"`          // 表示员工的个人邮箱，第三方仅通讯录应用可获取
	IsLeaderInDept   []int    `json:"is_leader_in_dept"` // 表示在所在的部门内是否为上级。；第三方仅通讯录应用可获取
	DirectLeader     string   `json:"direct_leader"`     // 与直接上级的关系；第三方仅通讯录应用可获取
	Avatar           string   `json:"avatar"`            // 头像url。 第三方仅通讯录应用可获取
	Telephone        string   `json:"telephone"`         // 座机。 第三方仅通讯录应用可获取
	Alias            string   `json:"alias"`             // 别名；第三方仅通讯录应用可获取
	Address          string   `json:"address"`           // 地址
	OpenUserID       string   `json:"open_userid"`       // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取。
	MainDepartment   int      `json:"main_department"`   // 主部门，仅当应用对主部门有查看权限时返回。
	Status           int      `json:"status"`            // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。已激活代表已激活企业微信或已关注微信插件（原企业号）。未激活代表既未激活企业微信又未关注微信插件（原企业号）。
	QrCode           string   `json:"qr_code"`           // 员工个人二维码，扫描可添加为外部联系人
	ExternalPosition string   `json:"external_position"` // 对外职务。当用户为企业内部员工时返回
	ExtAttr          struct { // 扩展属性，第三方仅通讯录应用可获取
		Attrs []UserInfoExtAttr `json:"attrs"`
	} `json:"extattr"`
}

type UserInfoExtAttr struct {
	Type int    `json:"type"`
	Name string `json:"name"`
	Text struct {
		Value string `json:"value"`
	} `json:"text"`
	Web struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"web"`
	Miniprogram struct {
		Appid    string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title    string `json:"title"`
	} `json:"miniprogram"`
}

type ExternalProfile struct {
	ExternalCorpName string `json:"external_corp_name"`
	WechatChannels   []struct {
		Nickname string `json:"nickname"`
		Status   int    `json:"status"`
	} `json:"wechat_channels"`
	ExternalAttr []UserInfoExtAttr `json:"external_attr"`
}

type GetUserListRequest struct {
	Cursor int `json:"cursor,omitempty"` // 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Limit  int `json:"limit,omitempty"`  // 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取默认值
}

type GetUserListResponse struct {
	wework.APIBaseResponse
	NextCursor int `json:"next_cursor"` // 分页游标，下次请求时填写以获取之后分页的记录。如果该字段返回空则表示已没有更多数据
	DeptUser   []struct {
		OpenUserID string `json:"open_userid"`
		UserID     string `json:"userid"`
		Department int    `json:"department"`
	} `json:"dept_list"`
}
