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
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	res, err := l.svcCtx.MerchantRpc.AddMerchant(l.ctx, &shop.AddMerchantReq{
		UserId:         int32(userId),
		MerchantName:   req.MerchantName,
		MerchantStatus: req.MerchantStatus,
		Linkname:       req.Linkname,
		Phone:          req.Phone,
	})
	if err != nil || !res.Ok {
		return nil, err
	}
	return
}
