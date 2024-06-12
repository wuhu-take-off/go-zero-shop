package raw_field

import (
	"fmt"
	"reflect"
	"strings"
)

// RawFieldNames
// 获取in结构体中gorm所映射的column字段
func RawFieldNames(in any) []string {
	out := make([]string, 0)

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		column := fi.Tag.Get("gorm")
		columns := strings.Split(column, ";")
		for _, s := range columns {
			if strings.HasPrefix(s, "column:") {
				out = append(out, s[7:])
			}
		}
	}
	return out
}

// UpdateFieldMap 获取更新字段键值对
func UpdateFieldMap(in any, change map[string]string) ([]string, []interface{}) {

	var sql []string
	var values []interface{}

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("ToMap only accepts structs; got %T", v))
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {

		// 获取字段的值
		fieldValue := v.Field(i)
		// 获取字段的类型
		//fieldType := fieldValue.Type()

		column := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
		c, ok := change[column]

		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() || fieldValue.Kind() != reflect.Ptr && ok {
			val := fieldValue
			if fieldValue.Kind() == reflect.Ptr {
				val = fieldValue.Elem()
			}
			if ok {
				sql = append(sql, c)
			} else {
				sql = append(sql, column)
			}
			switch val.Kind() {
			case reflect.Int32:
				values = append(values, val.Interface().(int32))
			case reflect.String:
				values = append(values, val.Interface().(string))
			}
			//fmt.Println(column, val)
		}
	}

	//fmt.Println(sql, values)
	return sql, values
}
