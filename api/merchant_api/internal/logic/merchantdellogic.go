package logic

import (
	"context"

	"TongChi_shop/api/merchant_api/internal/svc"
	"TongChi_shop/api/merchant_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantDelLogic {
	return &MerchantDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantDelLogic) MerchantDel(req *types.MerchantDelReq) error {
	// todo: add your logic here and delete this line

	return nil
}
