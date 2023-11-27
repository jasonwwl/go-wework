package wework

type GetProviderTokenResponse struct {
	APIBaseResponse
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int64  `json:"expires_in"`
}

type GetTokenResponse struct {
	APIBaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type AuthCorpAccessTokenResponse struct {
	APIBaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type SuiteTokenResponse struct {
	APIBaseResponse
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int64  `json:"expires_in"`
}
