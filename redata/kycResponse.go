package redata

// Identity 提交身份信息响应
type Identity struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Result  identityResult `json:"result"`
}

// identityResult 提交身份信息响应Result
type identityResult struct {
	Method int    `json:"method"`
	ID     int64  `json:"id"`
	Url    string `json:"url"`
}
