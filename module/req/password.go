package req

type SavePassword struct {
	UserName    string `json:"username" bson:"user_name"`
	PassWord    string `json:"password" bson:"password"`
	Description string `json:"description"  bson:"description"`
	Topic       string `json:"topic" bson:"topic"`
}
type QueryPasswd struct {
	Id string `form:"id"` //集合id
	Q  string `form:"q"`  //关键字
}
