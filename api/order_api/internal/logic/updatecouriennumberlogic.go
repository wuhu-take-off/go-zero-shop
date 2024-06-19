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

type UpdateCourienNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCourienNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourienNumberLogic {
	return &UpdateCourienNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourienNumberLogic) UpdateCourienNumber(req *types.UpdateCourierNumberReq) (resp *types.UpdateCourierNumberResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	number, err := l.svcCtx.OrderRpc.UpdateCourierNumber(l.ctx, &shop.UpdateCourierNumberReq{
		UserId:        int32(userId),
		OrderId:       req.OrderId,
		CourierNumber: req.CourierNumber,
	})
	if err != nil || !number.Ok {
		return nil, err
	}
	return
}
