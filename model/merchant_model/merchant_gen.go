package merchant_model

import (
	"TongChi_shop/model/user_model"
	"TongChi_shop/rpc/shop"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	usersMerchantSqlField = `
	merchant_infos.name AS merchant_name,
	merchant_infos.merchant_id AS merchant_id,
	merchant_infos.state AS status,
	merchant_infos.linkname AS linkname,
	merchant_infos.phone AS phone,
	user_infos.role AS Role
`
	UsersMerchantSqlFieldMap = map[string]string{
		"merchant_name":   "merchant_infos.name",
		"merchant_id":     "merchant_infos.merchant_id",
		"merchant_status": "merchant_infos.state",
	}
	MerchantInfosSqlFieldMap = map[string]string{
		"merchant_name": "name",
	}
)

type (
	merchantInfosModel interface {
		FindOneMerchant(userId int32) *MerchantInfos
		FindUsersMerchant(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersMerchant)
		UpdateMerchantInfos(merchantId int32, sql []string, values ...interface{}) error
		ConfirmUserRole(userId int32, merchantId int32) bool
		DelMerchantInfos(merchantId int32) error
		FindMerchantIdList(userId int32) []int32
	}
	defaultMerchantInfosModel struct {
		DB *gorm.DB
	}
	MerchantInfos struct {
		MerchantId int32     `gorm:"column:merchant_id;primaryKey"`
		UserId     int32     `gorm:"column:user_id"`
		Status     int32     `gorm:"column:state"`
		Name       string    `gorm:"column:name"`
		Linkname   string    `gorm:"column:linkname"`
		Phone      []byte    `gorm:"column:phone"`
		UpdateTime time.Time `gorm:"column:update_time"`
		DeleteTime time.Time `gorm:"column:delete_time"`
	}

	UsersMerchant struct {
		MerchantName   string `gorm:"column:merchant_name"`
		MerchantId     int32  `gorm:"column:merchant_id"`
		MerchantStatus int32  `gorm:"column:merchant_status"`
		Linkname       string `gorm:"column:linkname"`
		Phone          []byte `gorm:"column:phone"`
		Role           int32  `gorm:"column:role"`
	}
)

func (merchant *MerchantInfos) TableName() string {
	return "merchant_infos"
}
func (usersMerchant *UsersMerchant) ToMerchantListResp() *shop.MerchantList {
	return &shop.MerchantList{
		MerchantName:   usersMerchant.MerchantName,
		MerchantId:     usersMerchant.MerchantId,
		MerchantStatus: usersMerchant.MerchantStatus,
		Linkname:       usersMerchant.Linkname,
		Role:           usersMerchant.Role,
	}
}

func (m defaultMerchantInfosModel) FindOneMerchant(userId int32) *MerchantInfos {
	infos := &MerchantInfos{}
	m.DB.Where("user_id = ? AND delete_time IS NULL", userId).First(infos)
	return infos
}

// 统计用户能够管理的商铺数量
func (m *defaultMerchantInfosModel) CountUsersMerchant(userId ...interface{}) int32 {
	sql := ""
	if len(userId) == 1 {
		sql = "user_infos.user_id = ? AND "
	}
	sql = fmt.Sprintf(`
SELECT
 	COUNT(*)
FROM
	merchant_infos
	JOIN user_infos ON %s merchant_infos.user_id = user_infos.user_id
`, sql)
	var count int32
	m.DB.Raw(sql, userId...).Scan(&count)
	return count
}

// 获取用户能够管理的商铺数量和列表（分页）
func (m *defaultMerchantInfosModel) FindUsersMerchant(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersMerchant) {
	role := user_model.NewUserInfosModel(m.DB).UserRole(userId)
	switch role {
	case 1:
		return m.CountUsersMerchant(), m.admiUsersMerchant(page, limit, sql, values...)
	case 2:
		return m.CountUsersMerchant(userId), m.userUsersMerchant(userId, page, limit, sql, values...)
	default:
		return 0, nil
	}
}

// 获取普通商户管理的商铺列表
func (m *defaultMerchantInfosModel) userUsersMerchant(userId, page, limit int32, sql string, values ...interface{}) []UsersMerchant {
	offset := page*limit - limit
	var usersMerchant []UsersMerchant
	if values != nil {
		sql = "WHERE " + sql
	}

	values = append(values, nil)
	for i := len(values) - 1; i != 0; i-- {
		values[i] = values[i-1]
	}
	values[0] = userId

	values = append(values, limit, offset)
	sql = fmt.Sprintf("SELECT "+usersMerchantSqlField+`
FROM
	merchant_infos
	JOIN user_infos ON user_infos.user_id = ? AND merchant_infos.user_id = user_infos.user_id AND merchant_infos.delete_time IS NULL
	%s
LIMIT ? OFFSET ?; `, sql)
	m.DB.Raw(sql, values...).Find(&usersMerchant)
	return usersMerchant
}

// 获取管理员能够管理的商铺列表
func (m *defaultMerchantInfosModel) admiUsersMerchant(page, limit int32, sql string, values ...interface{}) []UsersMerchant {
	offset := page*limit - limit
	var usersMerchant []UsersMerchant
	if values != nil {
		sql = "WHERE " + sql
	}
	values = append(values, limit, offset)
	sql = fmt.Sprintf("SELECT "+usersMerchantSqlField+`
FROM
	merchant_infos
	JOIN user_infos ON merchant_infos.user_id = user_infos.user_id AND merchant_infos.delete_time IS NULL
	%s
LIMIT ? OFFSET ?; `, sql)
	m.DB.Raw(sql, values...).Find(&usersMerchant)
	return usersMerchant
}

// 验证用户是否能够管理某商铺
func (m *defaultMerchantInfosModel) ConfirmUserRole(userId int32, merchantId int32) bool {
	role := user_model.NewUserInfosModel(m.DB).UserRole(userId)
	switch role {
	case 1:
		return true
	case 2:
		var count int64
		m.DB.Model(&MerchantInfos{}).Where("user_id = ? AND merchant_id = ? ", userId, merchantId).Count(&count)
		return count != 0
	default:
		return false
	}
}

// 更新商铺信息
func (m *defaultMerchantInfosModel) UpdateMerchantInfos(merchantId int32, sql []string, values ...interface{}) error {
	val := make(map[string]interface{})
	for i := 0; i < len(sql); i++ {
		val[sql[i]] = values[i]
	}
	return m.DB.Model(&MerchantInfos{}).Where("merchant_id = ? ", merchantId).Updates(val).Error
}

// 删除商铺
func (m *defaultMerchantInfosModel) DelMerchantInfos(merchantId int32) error {
	return m.DB.Model(&MerchantInfos{}).Where("merchant_id = ?", merchantId).Update("delete_time", time.Now()).Error
}

// FindMerchantIdList 获取用户所管理的商铺ID列表
func (m defaultMerchantInfosModel) FindMerchantIdList(userId int32) []int32 {
	var merchantId []int32
	m.DB.Model(&MerchantInfos{}).Select("merchant_id").Where("user_id = ? AND delete_time IS NULL", userId).Find(&merchantId)
	return merchantId
}
