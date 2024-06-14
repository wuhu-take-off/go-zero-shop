package order_model

import (
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/model/user_model"
	"TongChi_shop/rpc/shop"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var (
	userOrderSqlField = `
	order_infos.order_id AS order_id,
	goods_infos.name AS goods_name,
	goods_specification_infos.size AS specification_size,
	order_infos.num AS order_num,
	recipient_infos.name AS recipient_name,
	recipient_infos.phone AS phone,
	recipient_infos.address AS address,
	goods_specification_infos.score AS score,
	order_infos.courier_number AS courier_number,
	order_infos.create_time AS create_time 
`
	UserOrderSqlFieldMap = map[string]string{
		"recipient_name":  "recipient_infos.name",
		"recipient_phone": "recipient_infos.phone",
		"order_time":      "order_infos.create_time",
		"goods_name":      "goods_infos.name",
		"order_id":        "order_infos.order_id",
	}
	userOrderSqlJoinField = `
FROM
	order_infos
	JOIN goods_specification_infos ON order_infos.specification_id = goods_specification_infos.specification_id
	JOIN goods_infos ON goods_specification_infos.goods_id = goods_infos.goods_id
	JOIN recipient_infos ON order_infos.recipient_id = recipient_infos.recipient_id
	
`
)

type (
	OrderInfosModel interface {
		FindUsersOrder(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersOrder)
		CheckUserRoleToUpCourierNumber(userId int32, orderId int64) bool
		UpdateCourierNumber(orderId int64, courierNumber string)
	}
	defaultOrderInfosModel struct {
		DB *gorm.DB
	}
	OrderInfos struct {
		OrderId         int64     `gorm:"column:order_id"`
		RecipientId     int32     `gorm:"column:recipient_id"`
		SpecificationId int32     `gorm:"column:specification_id"`
		Num             int32     `gorm:"column:num"`
		CourierNumber   string    `gorm:"column:courier_number"`
		CreateTime      time.Time `gorm:"column:create_time"`
	}
	UsersOrder struct {
		OrderId           int64     `gorm:"column:order_id"`
		GoodsName         string    `gorm:"column:goods_name"`
		SpecificationSize string    `gorm:"column:specification_size"`
		OrderNum          int32     `gorm:"column:order_num"`
		RecipientName     string    `gorm:"column:recipient_name"`
		Phone             string    `gorm:"column:phone"`
		Address           string    `gorm:"column:address"`
		Score             int32     `gorm:"column:score"`
		CourierNumber     string    `gorm:"column:courier_number"`
		CreateTime        time.Time `gorm:"column:create_time"`
	}
)

func (usersOrder *UsersOrder) ToOrderListResp() *shop.OrderList {
	return &shop.OrderList{
		CreateTime:        usersOrder.CreateTime.UnixNano(),
		OrderId:           usersOrder.OrderId,
		GoodsName:         usersOrder.GoodsName,
		SpecificationSize: usersOrder.SpecificationSize,
		OrderNum:          usersOrder.OrderNum,
		RecipientName:     usersOrder.RecipientName,
		Phone:             usersOrder.Phone,
		Address:           usersOrder.Address,
		Score:             usersOrder.Score,
		CourierNumber:     &usersOrder.CourierNumber,
	}
}

func (order *OrderInfos) TableName() string {
	return "order_infos"
}
func (m *defaultOrderInfosModel) CountUserOrder(userId ...interface{}) int32 {
	merchantInfosModel := merchant_model.NewMerchantInfosModel(m.DB)
	var merchantIdList []int32
	var sql string
	if userId != nil {
		merchantIdList = merchantInfosModel.FindMerchantIdList(userId[0].(int32))
		if merchantIdList == nil || len(merchantIdList) == 0 {
			return 0
		}
		sql = "WHERE goods_infos.merchant_id IN ? "
	} else {
		sql = ""
		merchantIdList = nil
	}

	sql = fmt.Sprintf(`
SELECT
 	COUNT(*)
FROM
	order_infos
	JOIN goods_specification_infos ON order_infos.specification_id = goods_specification_infos.specification_id
	JOIN goods_infos ON goods_specification_infos.goods_id = goods_infos.goods_id
	JOIN recipient_infos ON order_infos.recipient_id = recipient_infos.recipient_id
%s
`, sql)
	var count int32
	if merchantIdList != nil {
		m.DB.Raw(sql, merchantIdList).Scan(&count)
	} else {
		m.DB.Raw(sql).Scan(&count)
	}
	return count
}

// 返回用户可管理的订单信息
func (m *defaultOrderInfosModel) FindUsersOrder(userId, page, limit int32, sql string, values ...interface{}) (int32, []UsersOrder) {
	role := user_model.NewUserInfosModel(m.DB).UserRole(userId)

	switch role {
	case 1: //管理员能够看到所有的订单信息
		return m.CountUserOrder(), m.admiUsersOrder(page, limit, sql, values...)
	case 2: //商户可以看到自己商铺的订单信息
		return m.CountUserOrder(userId), m.userUsersOrder(userId, page, limit, sql, values...)
	}
	return 0, nil
}

func (m *defaultOrderInfosModel) admiUsersOrder(page, limit int32, sql string, values ...interface{}) []UsersOrder {
	offset := page*limit - limit
	var userOrder []UsersOrder
	if values != nil {
		sql = "WHERE " + sql
	}
	values = append(values, limit, offset)
	sql = fmt.Sprintf("SELECT "+userOrderSqlField+`
FROM
	order_infos
	JOIN goods_specification_infos ON order_infos.specification_id = goods_specification_infos.specification_id
	JOIN goods_infos ON goods_specification_infos.goods_id = goods_infos.goods_id
	JOIN recipient_infos ON order_infos.recipient_id = recipient_infos.recipient_id
	%s
LIMIT ? OFFSET ?; 
`, sql)
	m.DB.Raw(sql, values...).Find(&userOrder)
	return userOrder
}

func (m *defaultOrderInfosModel) userUsersOrder(userId, page, limit int32, sql string, values ...interface{}) []UsersOrder {
	offset := page*limit - limit
	var userOrder []UsersOrder
	if values != nil {
		sql = "WHERE " + sql + " AND goods_infos.merchant_id IN ? "
	} else {
		sql = "WHERE goods_infos.merchant_id IN ? "
	}

	merchantIdList := merchant_model.NewMerchantInfosModel(m.DB).FindMerchantIdList(userId)
	if len(merchantIdList) == 0 {
		return nil
	}

	values = append(values, merchantIdList, limit, offset)
	sql = fmt.Sprintf("SELECT "+userOrderSqlField+`
FROM
	order_infos
	JOIN goods_specification_infos ON order_infos.specification_id = goods_specification_infos.specification_id
	JOIN goods_infos ON goods_specification_infos.goods_id = goods_infos.goods_id
	JOIN recipient_infos ON order_infos.recipient_id = recipient_infos.recipient_id
	%s
LIMIT ? OFFSET ?; 
`, sql)
	m.DB.Raw(sql, values...).Find(&userOrder)
	return userOrder
}

func (m *defaultOrderInfosModel) CheckUserRoleToUpCourierNumber(userId int32, orderId int64) bool {
	role := user_model.NewUserInfosModel(m.DB).UserRole(userId)
	switch role {
	case 1:
		return true
	case 2:
		merchantIdList := merchant_model.NewMerchantInfosModel(m.DB).FindMerchantIdList(userId)
		if len(merchantIdList) == 0 {
			return false
		}
		sql := `
SELECT
 	COUNT(*)
FROM
	order_infos
	JOIN goods_specification_infos ON order_infos.order_id = ? AND order_infos.specification_id = goods_specification_infos.specification_id
	JOIN goods_infos ON goods_specification_infos.goods_id = goods_infos.goods_id
	JOIN recipient_infos ON order_infos.recipient_id = recipient_infos.recipient_id
WHERE goods_infos.merchant_id IN ? 
`
		var count int32
		m.DB.Raw(sql, orderId, merchantIdList).Scan(&count)
		return count != 0
	default:
		return false

	}
}

func (m *defaultOrderInfosModel) UpdateCourierNumber(orderId int64, courierNumber string) {
	m.DB.Model(&OrderInfos{}).Where("order_id = ?", orderId).Updates(map[string]interface{}{
		"courier_number": courierNumber,
	})
}
