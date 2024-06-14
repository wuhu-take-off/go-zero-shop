package merchantservelogic

import (
	"TongChi_shop/model/merchant_model"
	"context"
	"errors"

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
	merchantInfosModel := merchant_model.NewMerchantInfosModel(l.svcCtx.DB)
	if !merchantInfosModel.ConfirmUserRole(in.UserId, in.MerchantId) {
		return nil, errors.New("非法访问")
	}
	if err := merchantInfosModel.DelMerchantInfos(in.MerchantId); err != nil {
		return nil, errors.New("删除失败")
	}
	return &shop.DelMerchantResp{
		Ok: true,
	}, nil
}
