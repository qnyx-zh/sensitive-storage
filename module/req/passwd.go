package req

type SavePasswdReq struct {
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	Description string `json:"description"`
	Topic       string `json:"topic"`
}
type QueryPasswdReq struct {
	Id       string `form:"id"`       //集合id
	Q        string `form:"q"`        //关键字
	PageNum  int    `form:"pageNum"`  //分页: 第几页?从0开始
	PageSize int    `form:"pageSize"` //分页: 每页数据量
}
