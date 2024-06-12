package merchantservelogic

import (
	"context"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMerchantLogic {
	return &DelMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMerchantLogic) DelMerchant(in *shop.DelMerchantReq) (*shop.DelMerchantResp, error) {
	// todo: add your logic here and delete this line

	return &shop.DelMerchantResp{}, nil
}
