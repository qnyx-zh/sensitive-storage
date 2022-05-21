package resp

type PasswdInfosResp struct {
	//用户名
	Username string 
	//密码
	Password string
	//主题
	Topic string
	//id
	Id int
	//说明
	Description string
}

type PasswdInfoPageResp struct {
	result   []PasswdInfosResp
	pageNum  int
	pageSize int
	total    int
}
