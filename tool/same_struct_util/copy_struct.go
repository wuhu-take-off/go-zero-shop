package same_struct_util

import "reflect"

func CopyStruct(x, y interface{}) {
	vx := reflect.ValueOf(x).Elem()
	vy := reflect.ValueOf(y).Elem()

	for i := 0; i < vx.NumField(); i++ {
		fieldX := vx.Field(i)
		fieldY := vy.Field(i)
		fieldX.Set(fieldY)
	}
}
