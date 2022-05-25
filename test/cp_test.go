package test

import (
	"fmt"
	"sensitive-storage/module/ident"
	"sensitive-storage/util/copier"
	"testing"
)

func Test_cp(t *testing.T) {
	u := ident.User{
		UserName: "张三",
		Passwd:   "123456",
	}
	tar := ident.User{}
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("异常,%v", e)
		}
	}()
	copier.CopyVal(u, tar)
	panic("aaaa")
}