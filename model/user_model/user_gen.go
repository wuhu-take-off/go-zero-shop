package user_model

import (
	"TongChi_shop/model/specification_model"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/raw_field"
	"fmt"
	"gorm.io/gorm"
)

var (
	userGoodsFiledNames = raw_field.RawFieldNames(UsersGoods{})
	userGoodsSqlFile    = `		
		merchant_infos.name AS merchant_name,	
		merchant_infos.merchant_id AS merchant_id,
		goods_infos.goods_id,
		goods_infos.name AS goods_name,
		goods_infos.goods_type_id,
		goods_infos.state AS status,
		goods_infos.description AS goods_description
`
	UserGoodsSqlFileMap = map[string]string{
		"merchant_name":     "merchant_infos.name",
		"merchant_id":       "merchant_infos.merchant_id",
		"goods_id":          "goods_infos.goods_id",
		"goods_name":        "goods_infos.name",
		"goods_type_id":     "goods_infos.goods_type_id",
		"status":            "goods_infos.state",
		"goods_description": "goods_infos.description",
	}
)

type (
	userInfosModel interface {
		EmailExist(email string) bool
		UserRole(userId int32) int32
		FindUsersGoods(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersGoods)
	}

	defaultUserInfosModel struct {
		DB *gorm.DB
	}

	UserInfos struct {
		UserId int32  `gorm:"column:user_id;primaryKey"`
		Role   int32  `gorm:"column:role"`
		Email  string `gorm:"column:email"`
		Score  int32  `gorm:"column:score"`
	}

	// UsersGoods 用户的商品列表
	UsersGoods struct {
		MerchantName     string                                   `gorm:"column:merchant_name"`
		MerchantId       int32                                    `gorm:"column:merchant_id"`
		GoodsId          int32                                    `gorm:"column:goods_id"`
		GoodsName        string                                   `gorm:"column:goods_name"`
		GoodsTypeId      int32                                    `gorm:"column:goods_type_id"`
		GoodsState       int32                                    `gorm:"column:status"`
		GoodsDescription string                                   `gorm:"column:merchant_description"`
		Specifications   []specification_model.SpecificationInfos `gorm:"foreignKey:GoodsId;references:GoodsId"`
	}
)

func (users *UsersGoods) ToGoodsListResp() *shop.GoodsList {
	list := &shop.GoodsList{
		MerchantName:      users.MerchantName,
		GoodsId:           users.GoodsId,
		GoodsName:         users.GoodsName,
		GoodsTypeId:       users.GoodsTypeId,
		Status:            users.GoodsState,
		Describe:          users.GoodsDescription,
		SpecificationList: make([]*shop.SpecificationResp, 0, len(users.Specifications)),
	}
	for _, specification := range users.Specifications {
		list.SpecificationList = append(list.SpecificationList, specification.ToSpecificationResp())
	}
	return list
}

func (user *UserInfos) TableName() string {
	return "user_infos"
}

func (m *defaultUserInfosModel) EmailExist(email string) bool {
	var count int64
	m.DB.Model(&UserInfos{}).Where("email = ?", email).Count(&count)
	fmt.Println("email:", email, count)
	return count != 0
}

func (m *defaultUserInfosModel) UserRole(userId int32) int32 {
	var role int32
	m.DB.Model(&UserInfos{}).Select("role").Where("user_id = ?", userId).Scan(&role)
	return role
}

func (m *defaultUserInfosModel) CountUserGoods(userId ...interface{}) int32 {
	sql := fmt.Sprintf(`SELECT
		count(*)
		FROM
		user_infos 
		JOIN merchant_infos ON %s user_infos.user_id = merchant_infos.user_id
	 	JOIN goods_infos ON merchant_infos.merchant_id = goods_infos.merchant_id AND goods_infos.delete_time IS NULL
	`, func(userid ...interface{}) string {
		if len(userid) == 1 {
			return "user_infos.user_id = ? AND"
		}
		return ""
	}(userId...))
	var count int32
	m.DB.Raw(sql, userId...).Scan(&count)
	return count
}

func (m *defaultUserInfosModel) FindUsersGoods(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersGoods) {
	role := m.UserRole(userId)
	switch role {
	case 1: //管理员
		return m.CountUserGoods(), m.admiGoods(page, limit, sql, values...)
	case 2: //商户
		return m.CountUserGoods(userId), m.userGoods(userId, page, limit, sql, values...)
	default:
		return 0, nil
	}
}

func (m *defaultUserInfosModel) admiGoods(page, limit int32, sql string, values ...interface{}) []UsersGoods {
	offset := page*limit - limit
	var userGoods []UsersGoods
	if values != nil {
		sql = "WHERE " + sql
	}
	values = append(values, limit, offset)
	//sql = fmt.Sprintf("SELECT "+userGoodsSqlFile+`
	//FROM
	//	user_infos
	//	JOIN merchant_infos ON user_infos.user_id = merchant_infos.user_id
	// 	JOIN goods_infos ON merchant_infos.merchant_id = goods_infos.merchant_id AND goods_infos.delete_time IS NULL
	//	%s
	//	LIMIT ? OFFSET ?; `, sql)

	//修改为下面的sql语句，使得%s为sql成功占位
	baseSQL := "SELECT " + userGoodsSqlFile + `
    FROM
        user_infos
        JOIN merchant_infos ON user_infos.user_id = merchant_infos.user_id
        JOIN goods_infos ON merchant_infos.merchant_id = goods_infos.merchant_id AND goods_infos.delete_time IS NULL
        %s
    LIMIT ? OFFSET ?;`

	sql = fmt.Sprintf(baseSQL, sql)

	//获取用户所拥有的所有商户的商品信息
	m.DB.Raw(sql, values...).Find(&userGoods)
	return userGoods
}

func (m *defaultUserInfosModel) userGoods(userId, page, limit int32, sql string, values ...interface{}) []UsersGoods {
	offset := page*limit - limit
	var userGoods []UsersGoods
	if values != nil {
		sql = "WHERE " + sql
	}
	values = append(values, nil)
	for i := len(values) - 1; i != 0; i-- {
		values[i] = values[i-1]
	}
	values[0] = userId
	values = append(values, limit, offset)
	//sql = fmt.Sprintf("SELECT "+userGoodsSqlFile+`
	//FROM
	//	user_infos
	//	JOIN merchant_infos ON user_infos.user_id = ? AND user_infos.user_id = merchant_infos.user_id
	// 	JOIN goods_infos ON merchant_infos.merchant_id = goods_infos.merchant_id AND goods_infos.delete_time IS NULL
	//	%s
	//	LIMIT ? OFFSET ?; `, sql)

	//修改为下面的sql语句，使得%s为sql成功占位
	baseSQL := "SELECT " + userGoodsSqlFile + `
    FROM
        user_infos
        JOIN merchant_infos ON user_infos.user_id = merchant_infos.user_id
        JOIN goods_infos ON merchant_infos.merchant_id = goods_infos.merchant_id AND goods_infos.delete_time IS NULL
        %s
    LIMIT ? OFFSET ?;`

	sql = fmt.Sprintf(baseSQL, sql)

	//获取用户所拥有的所有商户的商品信息

	m.DB.Raw(sql, values...).Find(&userGoods)
	return userGoods
}
