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

type MerchantListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantListLogic {
	return &MerchantListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MerchantListLogic) MerchantList(req *types.MerchantListReq) (resp *types.MerchantListResp, err error) {
	value := l.ctx.Value("UserID")
	userId, err := json_number_transition.Transition(value)
	if err != nil {
		return nil, errors.New("非法访问")
	}

	merchantList, err := l.svcCtx.MerchantRpc.MerchantList(l.ctx, &shop.MerchantListReq{
		UserId:         int32(userId),
		MerchantStatus: req.MerchantStatus,
		MerchantName:   req.MerchantName,
		MerchantId:     req.MerchantId,
		Page:           req.Page,
		Limit:          req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.MerchantList
	for _, m := range merchantList.MerchantList {
		list = append(list, &types.MerchantList{
			MerchantName:   m.MerchantName,
			MerchantId:     m.MerchantId,
			MerchantStatus: m.MerchantStatus,
			Linkname:       m.Linkname,
			Role:           m.Role,
		})
	}

	return &types.MerchantListResp{
		Count:  merchantList.Count,
		Limit:  merchantList.Limit,
		Number: merchantList.Number,
		Page:   merchantList.Page,
		List:   list,
	}, nil
}
