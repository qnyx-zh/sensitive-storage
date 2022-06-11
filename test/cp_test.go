package test

import (
	"fmt"
	"reflect"
	"strings"
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
	//l := list.New()
	//l.PushFront(1111)
	//l.PushBack(2222)
	//fmt.Println(l.Front().Value)
	//a := &entity.User{BaseField: entity.BaseField{
	//	Id:        1,
	//	CreatedAt: 2,
	//	UpdatedAt: 3,
	//},
	//	Username: "zhangsan",
	//	Password: "123456",
	//}
	//of := reflect.TypeOf(a).Elem()
	//v := reflect.ValueOf(a).Elem()
	//for i := 0; i < of.NumField(); i++ {
	//	field := of.Field(i)
	//	value := v.Field(i)
	//	if field.Type.Kind() == reflect.Struct {
	//		numField := field.Type.NumField()
	//		for i := 0; i < numField; i++ {
	//			println(field.Type.Field(i).Name)
	//			name := value.FieldByName(field.Type.Field(i).Name)
	//			println(conv(name))
	//		}
	//	}
	//}
	//a := make([]int, 10)
	//b := append(a, 1)
	//z := append(a, 1)
	//k := a[0:10]
	//fmt.Print("aV---->")
	//for i := range a {
	//	fmt.Printf("%v   ", a[i])
	//}
	//fmt.Println("")
	//
	//fmt.Print("kV---->")
	//for i := range k {
	//	fmt.Printf("%v   ", k[i])
	//}
	//fmt.Println("")
	//k[9] = 100
	//fmt.Print("kV---->")
	//for i := range k {
	//	fmt.Printf("%v   ", k[i])
	//}
	//fmt.Println("")
	//fmt.Print("aV---->")
	//for i := range a {
	//	fmt.Printf("%v   ", a[i])
	//}
	//fmt.Println("")
	//for i := range a {
	//	print(a[i])
	//}
	//println("")
	//for i := range b {
	//	print(b[i])
	//}
	//println("")
	//c := 0
	//d := 0
	//e := 0
	//for i := range a {
	//	fmt.Print("a-->")
	//	fmt.Printf("第%d个，地址=%v ", c, &a[i])
	//	c++
	//}
	//println("")
	//b[1] = 1
	//for i := range b {
	//	fmt.Print("b-->")
	//	fmt.Printf("第%d个，地址=%v ", d, &b[i])
	//	d++
	//}
	//println("")
	//for i := range z {
	//	fmt.Print("z-->")
	//	fmt.Printf("第%d个，地址=%v ", e, &z[i])
	//	d++
	//}
	//test1(a)
	//println(DeferFunc1(1))
	//println(DeferFunc2(1))
	//println(DeferFunc3(1))
	s := "and a =b"
	if strings.HasPrefix(s, "and") {
		u := s[3:]
		println(u)
	}
}

func test1(a interface{}) {
	g := 0
	ints := a.([]int)
	ints[0] = 100
	println("")
	for i := range ints {
		fmt.Printf("第%d个，地址=%v ", g, &ints[i])
		g++
	}
}

func conv(v reflect.Value) interface{} {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		println(v.Int())
		return v.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint()
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Interface, reflect.Ptr:
		return v.Interface()
	}
	return nil
}
func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
