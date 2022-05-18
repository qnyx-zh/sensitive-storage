package req

type SavePasswdReq struct {
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
	Description string `json:"description"`
}
type QueryPasswdReq struct {
	Id       string //集合id
	q        string //关键字
	pageNum  int    //分页: 第几页?从0开始
	pageSize int    //分页: 每页数据量
}
