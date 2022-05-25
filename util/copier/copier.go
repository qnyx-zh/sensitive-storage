package copier

import (
	"errors"
	"log"
	"reflect"
)

// 结构体字段复制
func CopyVal(src interface{}, tar interface{}) (err error) {

	defer func() {
		if e := recover(); e != nil {
			log.Printf("%v", e)
			errors.New("异常了原因")
		}
	}()
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
	tarType, tarValue := reflect.TypeOf(tar), reflect.ValueOf(tar)
	if (tarType.Kind() != reflect.Ptr) || (tarType.Elem().Kind() != reflect.Struct) {
		_ = errors.New("目标必须为结构体指针")
	}
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}
	if srcType.Kind() != reflect.Struct {
		_ = errors.New("源对象必须为结构体")
	}
	tarType, tarValue = tarType.Elem(), tarValue.Elem()
	fieldQty := tarType.NumField()
	for i := 0; i < fieldQty; i++ {
		fieldName := tarType.Field(i)
		fieldValue := srcValue.FieldByName(fieldName.Name)
		if !fieldValue.IsValid() || fieldName.Type != fieldValue.Type() {
			continue
		}
		if tarValue.Field(i).CanSet() {
			tarValue.Field(i).Set(fieldValue)
		}
	}
	return nil
}
