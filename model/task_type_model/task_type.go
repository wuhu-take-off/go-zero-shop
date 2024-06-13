package task_type_model

import "gorm.io/gorm"

func NewTaskTypeModel(db *gorm.DB) TaskTypeModel {
	return &defaultTaskTypeModel{DB: db}
}
