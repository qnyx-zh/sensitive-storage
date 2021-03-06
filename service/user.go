package service

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"sensitive-storage/module/entity"
)

type userService struct {
}

func (u *userService) QueryByUsername(username string) any {
	var user entity.User
	first := Client.Where("username = ?", username).First(&user)
	if first.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return &user
}

func (u *userService) Save(user *entity.User) bool {
	Client.AutoMigrate()
	result := Client.Create(user)
	if result.Error != nil {
		log.Printf("%v", result.Error)
		return false
	}
	return true
}

func (u *userService) Query(e *entity.User) any {
	first := Client.First(e)
	if errors.Is(first.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	if first.Error == gorm.ErrRecordNotFound {
		return nil
	}
	return e
}
