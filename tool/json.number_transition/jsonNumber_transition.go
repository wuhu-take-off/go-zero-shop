package json_number_transition

import (
	"encoding/json"
	"errors"
	"strconv"
)

func Transition(data any) (int64, error) {
	var res int64
	switch v := data.(type) {
	case json.Number:
		if intValue, err := strconv.ParseInt(string(v), 10, 64); err == nil {
			res = intValue
		} else {
			//fmt.Println("无法将 json.Number 转换为 int32:", err)
			return 0, errors.New("转换失败")
		}
	default:
		//fmt.Println("anyValue 不是 json.Number 类型")
		return 0, errors.New("data不是json.Number类型")
	}
	return res, nil
}
