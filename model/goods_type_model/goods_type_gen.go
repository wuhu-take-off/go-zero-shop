package goods_type_model

import (
	"TongChi_shop/rpc/shop"
	"gorm.io/gorm"
)

type (
	GoodsTypeModel interface {
		FindGoodsTypes() []GoodsType
	}
	defaultGoodsTypeModel struct {
		DB *gorm.DB
	}
	GoodsType struct {
		GoodsTypeId   int32  `gorm:"column:goods_type_id;primaryKey"`
		GoodsTypeName string `gorm:"column:goods_type_name"`
	}
)

func (goodsType *GoodsType) TableName() string {
	return "goods_type"
}
func (goodsType *GoodsType) ToGoodsTypesResp() *shop.GoodsTypeList {
	return &shop.GoodsTypeList{
		GoodsTypeId:   goodsType.GoodsTypeId,
		GoodsTypeName: goodsType.GoodsTypeName,
	}
}

func (m *defaultGoodsTypeModel) FindGoodsTypes() []GoodsType {
	var goodsTypes []GoodsType
	m.DB.Find(&goodsTypes)
	return goodsTypes
}
