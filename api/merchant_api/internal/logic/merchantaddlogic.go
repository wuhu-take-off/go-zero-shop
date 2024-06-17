package logic

import (
	"context"

	"TongChi_shop/api/merchant_api/internal/svc"
	"TongChi_shop/api/merchant_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantAddLogic {
	return &MerchantAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantAddLogic) MerchantAdd(req *types.MerchantAddReq) (resp *types.MerchantAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
