package specification_model

import (
	"TongChi_shop/rpc/shop"
	"gorm.io/gorm"
)

type (
	SpecificationModel interface {
		FindSpecification(goodsId ...interface{}) []SpecificationInfos
		CreateSpecifications(specifications ...*SpecificationInfos) error
		UpSpecification(specificationId int32, value map[string]interface{}) error
		CheckSpecificationIdAndGoodsId(specificationId, goodsId int32) bool
	}

	defaultSpecificationModel struct {
		DB *gorm.DB
	}
	SpecificationInfos struct {
		SpecificationId  int32  `gorm:"column:specification_id;primaryKey"`
		GoodsId          int32  `gorm:"column:goods_id"`
		Size             string `gorm:"column:size"`
		DisplayInventory int32  `gorm:"column:display_inventory"`
		RealInventory    int32  `gorm:"column:real_inventory"`
		Score            int32  `gorm:"column:score"`
	}
)

func (specification *SpecificationInfos) TableName() string {
	return "goods_specification_infos"
}
func (specification *SpecificationInfos) ToSpecificationResp() *shop.SpecificationResp {
	return &shop.SpecificationResp{
		SpecificationId:  specification.SpecificationId,
		Size:             specification.Size,
		DisplayInventory: specification.DisplayInventory,
		RealInventory:    specification.RealInventory,
		Score:            specification.Score,
	}
}

// FindSpecification 获取商品id所对应的规格信息
func (m *defaultSpecificationModel) FindSpecification(goodsId ...interface{}) []SpecificationInfos {
	var specification []SpecificationInfos
	if len(goodsId) == 0 {
		return nil
	}
	m.DB.Where("goods_id IN ?", []interface{}{1, 2, 3}).Find(&specification)
	return specification
}
func (m *defaultSpecificationModel) CheckSpecificationIdAndGoodsId(specificationId, goodsId int32) bool {
	var count int64
	m.DB.Model(SpecificationInfos{}).Where("specification_id = ? AND goods_id = ?", specificationId, goodsId).Count(&count)
	return count != 0
}

func (m *defaultSpecificationModel) CreateSpecifications(specifications ...*SpecificationInfos) error {
	if len(specifications) == 0 {
		return nil
	}
	return m.DB.Create(specifications).Error
}

func (m *defaultSpecificationModel) UpSpecification(specificationId int32, value map[string]interface{}) error {
	if len(value) == 0 {
		return nil
	}
	return m.DB.Model(SpecificationInfos{}).Where("specification_id = ?", specificationId).Updates(value).Error
}
