package user_model

import "gorm.io/gorm"

type (
	UserInfosModel interface {
		userInfosModel
	}
)

func NewUserInfosModel(db *gorm.DB) UserInfosModel {
	return &defaultUserInfosModel{DB: db}
}
