package resp

type Resp struct {
	// 状态描述
	Status string `json:"status"`
	//异常信息
	ErrMsg string `json:"errMsg"`
	//状态码
	Code int `json:"code"`
	//返回数据
	Data any `json:"data"`
}
