package goods_model

import "gorm.io/gorm"

func NewGoodsInfosModel(db *gorm.DB) GoodsInfosModel {
	return &defaultGoodsInfosModel{
		DB: db,
	}
}
