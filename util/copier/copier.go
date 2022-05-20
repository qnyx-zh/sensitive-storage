package copier

import (
	"fmt"
	"reflect"
)

// 结构体字段复制
func CopyVal(source interface{}, target interface{}) {
	sType := reflect.TypeOf(source)
	sValue := reflect.ValueOf(&source).Elem()
	tValue := reflect.ValueOf(&target)
	for i := 0; i < sType.NumField(); i++ {
		filedName := sType.Field(i).Name
		filedValue := sValue.Field(i).Interface()
		if !tValue.FieldByName(filedName).CanSet() {
			continue
		}
		fmt.Printf("filedValue: %v\n", filedValue)
	}
}
