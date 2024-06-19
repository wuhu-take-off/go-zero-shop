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
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}
	merchant, err := l.svcCtx.MerchantRpc.UpdateMerchant(l.ctx, &shop.UpdateMerchantReq{
		UserId:         int32(userId),
		MerchantId:     req.MerchantId,
		MerchantName:   req.MerchantName,
		MerchantStatus: req.MerchantStatus,
		Linkname:       req.Linkname,
		Phone:          req.Phone,
	})
	if err != nil || !merchant.Ok {
		return nil, err
	}
	return
}
