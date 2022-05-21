package ident

type Passwd struct {
	Id          uint64 `bson:"id"`          //id
	UserId      int    `bson:"user_id"`     //用户id
	Username    string `bson:"user_name"`   //账户
	Password    string `bson:"password"`    //密码
	Description string `bson:"description"` //备注
	Topic       string `bson:"topic"`       //标题
}
