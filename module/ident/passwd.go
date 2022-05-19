package ident

type Passwd struct {
	Id          uint64 //id
	UserId      int    //用户id
	Username    string //账户
	Password    string //密码
	Description string //备注
	Topic       string //标题
}
