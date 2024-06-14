package orderservelogic

import (
	"TongChi_shop/model/order_model"
	"TongChi_shop/tool/raw_field"
	"context"
	"fmt"
	"strings"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderList 获取任务列表
func (l *OrderListLogic) OrderList(in *shop.OrderListReq) (*shop.OrderListResp, error) {
	fmt.Println(in)
	sql, val := raw_field.UpdateFieldMap(in, order_model.UserOrderSqlFieldMap)
	orderInfosModel := order_model.NewOrderInfosModel(l.svcCtx.DB)
	for i := 0; i < len(sql); i++ {
		sql[i] = sql[i] + " = ? "
	}
	count, orders := orderInfosModel.FindUsersOrder(in.UserId, in.Page, in.Limit, strings.Join(sql, " AND "), val...)
	var orderList []*shop.OrderList
	for i := 0; i < len(orders); i++ {
		orderList = append(orderList, orders[i].ToOrderListResp())
	}

	//fmt.Println(count, orders)
	return &shop.OrderListResp{
		OrderList: orderList,
		Count:     count,
		Limit:     in.Limit,
		Number:    (count + in.Limit - 1) / in.Limit,
		Page:      in.Page,
	}, nil
}
