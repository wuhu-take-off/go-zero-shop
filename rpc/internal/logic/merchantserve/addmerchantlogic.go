package merchantservelogic

import (
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/model/user_model"
	"TongChi_shop/tool/encrypt_AES"
	"context"
	"errors"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMerchantLogic {
	return &AddMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMerchantLogic) AddMerchant(in *shop.AddMerchantReq) (*shop.AddMerchantResp, error) {
	role := user_model.NewUserInfosModel(l.svcCtx.DB).UserRole(in.UserId)
	if role > 2 {
		return nil, errors.New("非法访问")
	}
	phone, err := encrypt_AES.AesEncrypt([]byte(in.Phone), []byte(l.svcCtx.Config.AES.Key))
	if err != nil {
		return nil, errors.New("添加失败")
	}
	if err := merchant_model.NewMerchantInfosModel(l.svcCtx.DB).AddMerchant(in.UserId, in.MerchantName, in.MerchantStatus, in.Linkname, phone); err != nil {
		return nil, errors.New("添加失败")
	}
	return &shop.AddMerchantResp{
		Ok: true,
	}, nil
}
