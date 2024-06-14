package merchant_model

import (
	"gorm.io/gorm"
	"time"
)

type (
	merchantInfosModel interface {
		FindOneMerchant(userId int32) *MerchantInfos
		FindMerchantIdList(userId int32) []int32
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

// FindMerchantIdList 获取用户所管理的商铺ID列表
func (m defaultMerchantInfosModel) FindMerchantIdList(userId int32) []int32 {
	var merchantId []int32
	m.DB.Model(&MerchantInfos{}).Select("merchant_id").Where("user_id = ? AND delete_time IS NULL", userId).Find(&merchantId)
	return merchantId
}
