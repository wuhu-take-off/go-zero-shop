package model

import (
	"TongChi_shop/model/user_model"
	"TongChi_shop/tool/db_init"
	"fmt"
	"testing"
)

func TestUserModelEmailExist(t *testing.T) {
	db := db_init.ConnMysql("mysql:112154ZhouM..@tcp(8.130.145.246)/TongChi_shop?charset=utf8mb4&parseTime=True&loc=Local")
	UserModel := user_model.NewUserInfosModel(db)
	count, goods := UserModel.FindUsersGoods(1, 1, 3, "goods_infos.state = ?", 1)

	fmt.Println(count, goods)
}

func Test(t *testing.T) {
	//db := db_init.ConnMysql("mysql:112154ZhouM..@tcp(8.130.145.246)/TongChi_shop?charset=utf8mb4&parseTime=True&loc=Local")
	//model := goods_model.NewGoodsInfosModel(db)
	//err := model.DelGoodsInfos(20)
	//fmt.Println(err)
	run(nil)
}
func run(mp map[string]string) {
	if s, ok := mp["20"]; ok {
		fmt.Println(s, ok)
	} else {
		fmt.Println(s, ok)
	}
}
