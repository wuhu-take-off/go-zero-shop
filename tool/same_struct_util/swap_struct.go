package same_struct_util

import "reflect"

//交换两个相同结构的结构体

func SwapStructs(x, y interface{}) {
	vx := reflect.ValueOf(x).Elem()
	vy := reflect.ValueOf(y).Elem()

	for i := 0; i < vx.NumField(); i++ {
		fieldX := vx.Field(i)
		fieldY := vy.Field(i)
		tmp := reflect.New(fieldX.Type()).Elem()
		tmp.Set(fieldX)
		fieldX.Set(fieldY)
		fieldY.Set(tmp)
	}
}
