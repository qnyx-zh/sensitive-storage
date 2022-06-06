package service

import (
	"sensitive-storage/module/entity"
)

type passwordService struct {
}

func (p *passwordService) Save(passwd *entity.Password, userId uint) bool {
	return true
}
func (p *passwordService) QueryPasswordById(id uint, userId uint) entity.Password {
	var result entity.Password
	return result
}
