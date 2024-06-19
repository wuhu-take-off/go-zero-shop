package logic

import (
	"TongChi_shop/rpc/shop"
	json_number_transition "TongChi_shop/tool/json.number_transition"
	"context"
	"errors"

	"TongChi_shop/api/order_api/internal/svc"
	"TongChi_shop/api/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req *types.OrderListReq) (resp *types.OrderListResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	orderList, err := l.svcCtx.OrderRpc.OrderList(l.ctx, &shop.OrderListReq{
		UserId:         int32(userId),
		RecipientName:  req.RecipientName,
		RecipientPhone: req.RecipientPhone,
		OrderTime:      req.OrderTime,
		GoodsName:      req.GoodsName,
		OrderId:        req.OrderId,
		Page:           req.Page,
		Limit:          req.Limit,
	})
	if err != nil {
		return nil, err
	}
	var list []*types.OrderList
	for _, o := range orderList.OrderList {
		list = append(list, &types.OrderList{
			CreateTime:        o.CreateTime,
			OrderId:           o.OrderId,
			GoodsName:         o.GoodsName,
			SpecificationSize: o.SpecificationSize,
			OrderNum:          o.OrderNum,
			RecipientName:     o.RecipientName,
			Phone:             o.Phone,
			Address:           o.Address,
			Score:             o.Score,
			CourierNumber:     o.CourierNumber,
		})
	}
	return &types.OrderListResp{
		Count:  orderList.Count,
		Limit:  orderList.Limit,
		Number: orderList.Number,
		Page:   orderList.Page,
		List:   list,
	}, nil
}
