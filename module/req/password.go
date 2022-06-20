package req

type SavePassword struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Topic       string `json:"topic"`
}
type QueryPasswd struct {
	Q        string `form:"q"`                           // 关键字
	PageNum  *uint  `form:"pageNum" binding:"required"`  // 分页: 第几页?从0开始
	PageSize *uint  `form:"pageSize" binding:"required"` // 分页: 每页数据量
}

func (sp *SavePassword) ToUpdate() bool {
	if sp.Id > 0 {
		return true
	}
	return false
}
