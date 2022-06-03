package test

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_cp(t *testing.T) {
	//u := ident.User{
	//	UserName: "张三",
	//	Passwd:   "123456",
	//}
	//tar := ident.User{}
	//defer func() {
	//	if e := recover(); e != nil {
	//		fmt.Printf("异常,%v", e)
	//	}
	//}()
	//copier.CopyVal(u, tar)
	l := list.New()
	l.PushFront(1111)
	l.PushBack(2222)
	fmt.Println(l.Front().Value)
}
