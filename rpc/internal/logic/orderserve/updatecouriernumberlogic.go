package orderservelogic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &shop.UpdateCourierNumberResp{}, nil
}
