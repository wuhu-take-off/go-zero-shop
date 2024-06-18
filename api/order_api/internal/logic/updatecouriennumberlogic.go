package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
