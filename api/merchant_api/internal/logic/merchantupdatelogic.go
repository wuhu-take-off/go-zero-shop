package logic

import (
	"context"

	"TongChi_shop/api/merchant_api/internal/svc"
	"TongChi_shop/api/merchant_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantUpdateLogic {
	return &MerchantUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantUpdateLogic) MerchantUpdate(req *types.MerchantUpDateReq) (resp *types.MerchantUpDateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
