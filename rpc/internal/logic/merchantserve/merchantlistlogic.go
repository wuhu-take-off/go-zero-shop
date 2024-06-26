package merchantservelogic

import (
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/tool/encrypt_AES"
	"TongChi_shop/tool/raw_field"
	"context"
	"strings"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type MerchantListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMerchantListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MerchantListLogic {
	return &MerchantListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MerchantListLogic) MerchantList(in *shop.MerchantListReq) (*shop.MerchantListResp, error) {
	sql, val := raw_field.UpdateFieldMap(in, merchant_model.UsersMerchantSqlFieldMap)
	for i := 0; i < len(sql); i++ {
		sql[i] = sql[i] + " = ? "
	}
	merchantInfosModel := merchant_model.NewMerchantInfosModel(l.svcCtx.DB)
	count, merchants := merchantInfosModel.FindUsersMerchant(in.UserId, in.Page, in.Limit, strings.Join(sql, " AND "), val...)

	var merchantList []*shop.MerchantList
	for i := 0; i < len(merchants); i++ {
		resp := merchants[i].ToMerchantListResp()
		//解密敏感信息
		tmp, err := encrypt_AES.AesDecrypt(merchants[i].Phone, []byte(l.svcCtx.Config.AES.Key))
		if err != nil {
			continue
		}
		resp.Phone = string(tmp)
		merchantList = append(merchantList, resp)
	}
	return &shop.MerchantListResp{
		MerchantList: merchantList,
		Count:        count,
		Limit:        in.Limit,
		Number:       (count + in.Limit) / in.Limit,
		Page:         in.Page,
	}, nil
}
