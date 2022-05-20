package test

import (
	"fmt"
	"reflect"
	"sensitive-storage/module/req"
	"testing"
)

func Test_copy(t *testing.T) {
	save1 := req.SavePasswdReq{
		UserName: "zzz",
		PassWord: "123456",
	}
	printXXXX(save1)

}
func printXXXX(save1 interface{}){
	
	sT := reflect.TypeOf(save1)

	v := reflect.New(sT.Elem())
	fmt.Printf("v: %v\n", v)
	sV := reflect.ValueOf(&save1).Elem()
	for i := 0; i < sT.NumField(); i++ {
		s := sT.Field(i).Name
		sV.FieldByName(s).Set(reflect.ValueOf("aaaa"))
	}
	fmt.Println(save1)
}