package order_model

import "gorm.io/gorm"

func NewOrderInfosModel(db *gorm.DB) OrderInfosModel {
	return &defaultOrderInfosModel{
		DB: db,
	}
}
