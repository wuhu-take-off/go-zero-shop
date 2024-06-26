package task_model

import "gorm.io/gorm"

var (
	taskUpdateField = []string{"score"}
)

type (
	TaskInfosModel interface {
		UpdateTaskScore(taskId int32, score int32) error
	}
	defaultTaskInfosModel struct {
		DB *gorm.DB
	}
	TaskInfos struct {
		TaskId      int32  `gorm:"column:task_id;primaryKey"`
		MerchantId  int32  `gorm:"column:merchant_id"`
		Name        string `gorm:"column:name"`
		TaskTypeId  int32  `gorm:"column:task_type_id"`
		Status      int32  `gorm:"column:state"`
		Description string `gorm:"column:description"`
		Score       int32  `gorm:"column:score"`
	}
)

func (task *TaskInfos) TableName() string {
	return "task_infos"
}

func (m *defaultTaskInfosModel) UpdateTaskScore(taskId int32, score int32) error {
	return m.DB.Model(&TaskInfos{}).Where("task_id = ?", taskId).Updates(map[string]interface{}{
		"score": score,
	}).Error
}
