package orderservelogic

import (
	"TongChi_shop/model/order_model"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourierNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourierNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourierNumberLogic {
	return &UpdateCourierNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新任务信息
func (l *UpdateCourierNumberLogic) UpdateCourierNumber(in *shop.UpdateCourierNumberReq) (*shop.UpdateCourierNumberResp, error) {
	orderModel := order_model.NewOrderInfosModel(l.svcCtx.DB)
	if !orderModel.CheckUserRoleToUpCourierNumber(in.UserId, in.OrderId) {
		return nil, errors.New("非法访问")
	}
	orderModel.UpdateCourierNumber(in.OrderId, in.CourierNumber)
	return &shop.UpdateCourierNumberResp{
		Ok: true,
	}, nil
}
