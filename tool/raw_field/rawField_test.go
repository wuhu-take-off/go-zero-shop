package raw_field

import (
	"testing"
)

type GoodsAddReq struct {
	GoodsName   string  `json:"goods_name" maxValue:"20" minValue:"1"`
	GoodsType   int32   `json:"goods_type" maxValue:"2147483647" minValue:"1"`
	Status      int32   `json:"status" maxValue:"1" minValue:"0"`
	Description *string `json:"description,optional" maxValue:"100" minValue:"1"`
}

func TestRawFieldNames(t *testing.T) {
	CheckMaxMin(GoodsAddReq{
		GoodsName:   "aefa",
		GoodsType:   3123,
		Status:      20,
		Description: nil,
	})
}
