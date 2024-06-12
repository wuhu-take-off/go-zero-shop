package specification_model

import "gorm.io/gorm"

func NewSpecificationModel(db *gorm.DB) SpecificationModel {
	return &defaultSpecificationModel{DB: db}
}
