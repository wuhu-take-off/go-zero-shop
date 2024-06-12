package merchant_model

import "gorm.io/gorm"

type (
	MerchantInfosModel interface {
		merchantInfosModel
	}
)

func NewMerchantInfosModel(db *gorm.DB) MerchantInfosModel {
	return &defaultMerchantInfosModel{DB: db}
}
