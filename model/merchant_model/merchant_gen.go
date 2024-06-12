package merchant_model

import (
	"gorm.io/gorm"
	"time"
)

type (
	merchantInfosModel interface {
		FindOneMerchant(userId int32) *MerchantInfos
	}
	defaultMerchantInfosModel struct {
		DB *gorm.DB
	}
	MerchantInfos struct {
		MerchantId int32     `gorm:"column:merchant_id;primaryKey"`
		UserId     int32     `gorm:"column:user_id"`
		Status     int32     `gorm:"column:state"`
		Name       string    `gorm:"column:name"`
		Linkname   string    `gorm:"column:linkname"`
		Phone      string    `gorm:"column:phone"`
		UpdateTime time.Time `gorm:"column:update_time"`
		DeleteTime time.Time `gorm:"column:delete_time"`
	}
)

func (merchant *MerchantInfos) TableName() string {
	return "merchant_infos"
}
func (m defaultMerchantInfosModel) FindOneMerchant(userId int32) *MerchantInfos {
	infos := &MerchantInfos{}
	m.DB.Where("user_id = ? AND delete_time IS NULL", userId).First(infos)
	return infos
}
