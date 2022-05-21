package req

type LoginReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RegisterReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
