package imgs_modelr

import "gorm.io/gorm"

func NewImgsModel(db *gorm.DB) ImgsModel {
	return &defaultImgsModel{DB: db}
}
