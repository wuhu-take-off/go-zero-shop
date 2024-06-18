package merchantservelogic

import (
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/encrypt_AES"
	"TongChi_shop/tool/raw_field"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMerchantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMerchantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMerchantLogic {
	return &UpdateMerchantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMerchantLogic) UpdateMerchant(in *shop.UpdateMerchantReq) (*shop.UpdateMerchantResp, error) {
	merchantModel := merchant_model.NewMerchantInfosModel(l.svcCtx.DB)
	//验证身份
	if !merchantModel.ConfirmUserRole(in.UserId, in.MerchantId) {
		return nil, errors.New("非法访问")
	}

	//对电话号码进行加密
	if in.Phone != nil {
		aes, err := encrypt_AES.AesEncrypt([]byte(*in.Phone), []byte(l.svcCtx.Config.AES.Key))
		if err != nil {
			return nil, errors.New("更新失败")
		}
		phone := string(aes)
		in.Phone = &phone
	}

	//status字段单独处理
	if in.MerchantStatus != nil {

		in.MerchantStatus = nil
	}
	sql, val := raw_field.UpdateFieldMap(in, merchant_model.MerchantInfosSqlFieldMap)
	if err := merchantModel.UpdateMerchantInfos(in.MerchantId, sql, val...); err != nil {
		return nil, errors.New("更新失败")
	}

	return &shop.UpdateMerchantResp{
		Ok: true,
	}, nil
}
