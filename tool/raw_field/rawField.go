package raw_field

import (
	"fmt"
	"reflect"
	"strconv"
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
			case reflect.Int64:
				values = append(values, val.Interface().(int64))
			default:
				sql = sql[:len(sql)-1]
			}
			//fmt.Println(column, val)
		}
	}

	//fmt.Println(sql, values)
	return sql, values
}

// CheckMaxMin
// 根据tag中的maxValue和minValue判断数据是否合格,数值类型将会判断值,字符类型将会判断长度
func CheckMaxMin(in any) bool {
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
		maxTmp := fi.Tag.Get("maxValue")
		minTmp := fi.Tag.Get("minValue")
		if maxTmp == "" || minTmp == "" {
			continue
		}

		maxValue, _ := strconv.ParseInt(maxTmp, 10, 64)
		minValue, _ := strconv.ParseInt(minTmp, 10, 64)
		fmt.Println(maxValue, minValue)
		fieldValue := v.Field(i)
		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}
		switch fieldValue.Kind() {
		case reflect.Int32:
			val := fieldValue.Interface().(int32)
			if val < int32(minValue) || val > int32(maxValue) {
				return false
			}
		case reflect.String:
			val := fieldValue.Interface().(string)
			if len(val) < int(minValue) || len(val) > int(maxValue) {
				return false
			}
		case reflect.Int64:
			val := fieldValue.Interface().(int64)
			if val < minValue || val > maxValue {
				return false
			}
		}
	}
	return true
}
