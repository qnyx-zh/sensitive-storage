package test

import (
	"sensitive-storage/module/entity"
	"sensitive-storage/service"
	"testing"
)

func TestUtil(t *testing.T) {
	r := &entity.User{}
	service.GeneralDB.LambdaQuery().Eq("id", 1).One(r)
	print(r)
}
