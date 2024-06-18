package goods_type_model

import "gorm.io/gorm"

func NewGoodsTypeModel(db *gorm.DB) GoodsTypeModel {
	return &defaultGoodsTypeModel{DB: db}
}
