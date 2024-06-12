package goodsservelogic

import (
	"TongChi_shop/model/goods_model"
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/model/specification_model"
	"TongChi_shop/model/user_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"TongChi_shop/tool/raw_field"
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsUpdateLogic {
	return &GoodsUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *GoodsUpdateLogic) GoodsUpdate(in *shop.GoodsUpdateReq) (*shop.GoodsUpdateResp, error) {
	// todo: add your logic here and delete this line
	//权限验证
	role := user_model.NewUserInfosModel(l.svcCtx.DB).UserRole(in.UserId)
	switch role {
	case 1:
		{
		}
	case 2:
		merchantId := merchant_model.NewMerchantInfosModel(l.svcCtx.DB).FindOneMerchant(in.UserId).MerchantId
		if !goods_model.NewGoodsInfosModel(l.svcCtx.DB).ExistMerchantIdAndGoodsId(merchantId, in.GoodsId) {
			return nil, errors.New("禁止访问")
		}
	default:
		return nil, errors.New("禁止访问")
	}

	sql, val := raw_field.UpdateFieldMap(in, map[string]string{
		"status":     "state",
		"goods_name": "name",
		"describe":   "description",
	})
	//fmt.Println(sql, val)
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {

		//修改商品信息
		goodsInfosModel := goods_model.NewGoodsInfosModel(tx)
		value := make(map[string]interface{})
		for i := 0; i < len(sql); i++ {
			value[sql[i]] = val[i]
		}
		if err := goodsInfosModel.UpGoodsInfos(in.GoodsId, value); err != nil {
			return err
		}

		//修改规格信息
		specificationModel := specification_model.NewSpecificationModel(tx)
		for _, resp := range in.SpecificationList {
			if resp.SpecificationId == 0 {
				if err := specificationModel.CreateSpecifications(&specification_model.SpecificationInfos{
					GoodsId:          in.GoodsId,
					Size:             resp.Size,
					DisplayInventory: resp.DisplayInventory,
					RealInventory:    resp.RealInventory,
					Score:            resp.Score,
				}); err != nil {
					return err
				}
			} else {
				if !specificationModel.CheckSpecificationIdAndGoodsId(resp.SpecificationId, in.GoodsId) {
					continue
				}

				sql, val = raw_field.UpdateFieldMap(resp, map[string]string{
					"size":              "size",
					"display_inventory": "display_inventory",
					"real_inventory":    "real_inventory",
					"score":             "score",
				})
				//fmt.Println(sql, val)
				value = make(map[string]interface{})
				for i := 0; i < len(sql); i++ {
					value[sql[i]] = val[i]
				}
				if err := specificationModel.UpSpecification(resp.SpecificationId, value); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &shop.GoodsUpdateResp{
		OK: true,
	}, nil
}
