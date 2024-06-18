package imgs_modelr

import "gorm.io/gorm"

type (
	ImgsModel interface {
		Add(goodsId int32, imgName string) (int32, error)
		FindImgs(goodsId int32) []Imgs
	}
	defaultImgsModel struct {
		DB *gorm.DB
	}
	Imgs struct {
		ImgId   int32  `gorm:"column:img_id;primaryKey"`
		GoodsId int32  `gorm:"column:goods_id"`
		ImgName string `gorm:"column:img_name"`
	}
)

func (i *Imgs) TableName() string {
	return "imgs"
}

func (m *defaultImgsModel) Add(goodsId int32, imgName string) (int32, error) {
	img := &Imgs{
		GoodsId: goodsId,
		ImgName: imgName,
	}
	err := m.DB.Create(img).Error
	return img.ImgId, err
}

func (m *defaultImgsModel) FindImgs(goodsId int32) []Imgs {
	var imgs []Imgs
	m.DB.Model(&Imgs{}).Where("goods_id = ?", goodsId).Find(&imgs)
	return imgs
}
