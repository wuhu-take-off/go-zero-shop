package task_type_model

import (
	"TongChi_shop/rpc/shop"
	"gorm.io/gorm"
)

type (
	TaskTypeModel interface {
		FindTaskTypes() []TaskType
	}
	defaultTaskTypeModel struct {
		DB *gorm.DB
	}
	TaskType struct {
		TaskTypeId   int32  `gorm:"column:goods_type_id;primaryKey"`
		TaskTypeName string `gorm:"column:goods_type_name"`
	}
)

func (taskType *TaskType) TableName() string {
	return "goods_type"
}
func (taskType *TaskType) ToTaskTypesResp() *shop.TaskTypeList {
	return &shop.TaskTypeList{
		TaskTypeId:   taskType.TaskTypeId,
		TaskTypeName: taskType.TaskTypeName,
	}
}

func (m *defaultTaskTypeModel) FindTaskTypes() []TaskType {
	var taskTypes []TaskType
	m.DB.Find(&taskTypes)
	return taskTypes
}
