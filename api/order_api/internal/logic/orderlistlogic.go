package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.OrderListResp{
		Count:  10,
		Limit:  20,
		Number: 390,
		Page:   34,
		List:   nil,
	}, nil
}
