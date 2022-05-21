package ident

type Page struct {
	Current int         //当前页
	Size    int         //每页大小
	Total   int         //总数
	Data    interface{} //数据
}

func (p *Page) NewPage(cur, size, total int, data interface{}) *Page {
	return &Page{
		Current: cur,
		Size:    size,
		Total:   total,
		Data:    data,
	}
}
