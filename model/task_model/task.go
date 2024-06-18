package task_model

import "gorm.io/gorm"

func NewTaskInfosModel(db *gorm.DB) TaskInfosModel {
	return &defaultTaskInfosModel{DB: db}
}
