package my_reflect

import (
	"errors"
	"reflect"
	"strings"
)

func IsBlank(i interface{}) bool {
	value := reflect.ValueOf(i)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// 获取列名
func getColumnName(t reflect.StructField) string {
	gorm := t.Tag.Get("gorm")
	tags := strings.Split(gorm, ";")
	for _, v := range tags {
		if strings.Contains(v, "column:") {
			return strings.Split(v, ":")[1]
		}
	}
	return t.Name
}

func GetNameAndValue(t interface{}) map[string]interface{} {
	tType := reflect.TypeOf(t)
	if (tType.Kind() != reflect.Ptr) || (tType.Elem().Kind() != reflect.Struct) {
		errors.New("实体必须为结构体/结构体数组指针")
		return nil
	}
	tType = tType.Elem()
	tValue := reflect.ValueOf(t).Elem()
	m := make(map[string]interface{})
	fieldNum := tType.NumField()
	for i := 0; i < fieldNum; i++ {
		field := tType.Field(i)
		childV := tValue.Field(i)
		if field.Type.Kind() == reflect.Struct {
			for j := 0; j < field.Type.NumField(); j++ {
				childField := field.Type.Field(j)
				childValueField := childV.FieldByName(childField.Name)
				m[getColumnName(childField)] = childValueField.Interface()
			}
			continue
		}
		m[getColumnName(field)] = tValue.FieldByName(field.Name)
	}
	return m
}
