package service

import (
	"gorm.io/gorm"
	"log"
	"sensitive-storage/module/entity"
)

type passwordService struct {
}

func (p *passwordService) SearchPasswdList(id uint, q string, page *entity.Page) *entity.Page {
	var result []entity.Password
	find := Client.Where("id = ? and Topic like ?", id, "%"+q+"%").Offset((page.Cur - 1) * page.Size).Limit(page.Size).Find(&result)
	if find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		log.Println("数据不存在")
		return page
	}
	var count int64
	find.Count(&count)
	page.Data = result
	page.Total = count
	return page
}
