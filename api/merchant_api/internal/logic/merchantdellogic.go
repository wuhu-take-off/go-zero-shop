package logic

import (
	"TongChi_shop/rpc/shop"
	json_number_transition "TongChi_shop/tool/json.number_transition"
	"context"
	"errors"

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
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return errors.New("非法访问")
	}
	merchant, err := l.svcCtx.MerchantRpc.DelMerchant(l.ctx, &shop.DelMerchantReq{
		UserId:     int32(userId),
		MerchantId: req.MerchantId,
	})
	if err != nil || !merchant.Ok {
		return err
	}
	return nil
}
