package goods_model

import (
	"gorm.io/gorm"
	"time"
)

var (
	//create的忽略字段
	goodsInfosListOmitList = "update_time,delete_time"
)

type (
	GoodsInfosModel interface {
		CreateGoodsInfos(goods *GoodsInfos) error
		DelGoodsInfos(goodsId int32) error
		UpGoodsInfos(goodsId int32, values map[string]interface{}) error
		ExistMerchantIdAndGoodsId(merchantId, goodsId int32) bool
	}
	defaultGoodsInfosModel struct {
		DB *gorm.DB
	}
	GoodsInfos struct {
		GoodsId     int32     `gorm:"column:goods_id;primaryKey"`
		MerchantId  int32     `gorm:"column:merchant_id"`
		Name        string    `gorm:"column:name"`
		GoodsTypeId int32     `gorm:"column:goods_type_id"`
		Status      int32     `gorm:"column:state"`
		Description string    `gorm:"column:description"`
		Img         string    `gorm:"column:img"`
		UpdateTime  time.Time `gorm:"column:update_time"`
		DeleteTime  time.Time `gorm:"column:delete_time"`
	}
)

func (goods *GoodsInfos) TableName() string {
	return "goods_infos"
}
func (m *defaultGoodsInfosModel) CreateGoodsInfos(goods *GoodsInfos) error {
	return m.DB.Omit(goodsInfosListOmitList).Create(goods).Error
}

func (m *defaultGoodsInfosModel) DelGoodsInfos(goodsId int32) error {
	return m.DB.Model(GoodsInfos{}).Where("goods_id = ?", goodsId).Updates(map[string]interface{}{"delete_time": time.Now()}).Error
}
func (m *defaultGoodsInfosModel) UpGoodsInfos(goodsId int32, values map[string]interface{}) error {
	if len(values) == 0 {
		return nil
	}
	return m.DB.Model(GoodsInfos{}).Where("goods_id = ?", goodsId).Updates(values).Error
}

func (m *defaultGoodsInfosModel) ExistMerchantIdAndGoodsId(merchantId, goodsId int32) bool {
	var count int64
	m.DB.Model(GoodsInfos{}).Where("goods_id = ? AND merchant_id = ?", goodsId, merchantId).Count(&count)
	return count != 0
}
