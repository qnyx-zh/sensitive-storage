package resp

type PasswdInfosResp struct {
	Id          uint64 `bson:"id" json:"id"`                  //id
	UserId      int    `bson:"user_id" json:"userId"`          //用户
	Username    string `bson:"user_name" json:"username"`      //账户
	Password    string `bson:"password" json:"password"`       //密码
	Description string `bson:"description" json:"description"` //备注
	Topic       string `bson:"topic" json:"topic"`             //标题
}
type LoginResp struct {
	Token string `json:"token"` //token
}

type Passwd struct {
	Passwd []PasswdInfosResp `json:"passwds"`
}
