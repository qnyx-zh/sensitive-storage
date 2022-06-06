package resp

type PasswdInfos struct {
	Id          string `json:"id"`          //id
	UserId      string `json:"userId"`      //所属人id
	Username    string `json:"username"`    //账户
	Password    string `json:"password"`    //密码
	Description string `json:"description"` //备注
	Topic       string `json:"topic"`       //标题
}
type Login struct {
	Token string `json:"token"` //token
}

type Passwd struct {
	Password []PasswdInfos `json:"passwds"`
}
