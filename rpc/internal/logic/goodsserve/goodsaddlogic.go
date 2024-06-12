package goodsservelogic

import (
	"TongChi_shop/model/goods_model"
	"TongChi_shop/model/merchant_model"
	"TongChi_shop/model/specification_model"
	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsAddLogic {
	return &GoodsAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GoodsAdd
// 添加商品
func (l *GoodsAddLogic) GoodsAdd(in *shop.GoodsAddReq) (*shop.GoodsAddResp, error) {
	// todo: add your logic here and delete this line

	if in.SpecificationList == nil {
		return nil, errors.New("规格不能为空")
	}
	merchantModel := merchant_model.NewMerchantInfosModel(l.svcCtx.DB)
	merchant := merchantModel.FindOneMerchant(in.UserId)
	if merchant.MerchantId == 0 {
		return nil, errors.New("该用户没有创建商户")
	}

	//通过事务进行添加操作
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//添加商品
		goodsModel := goods_model.NewGoodsInfosModel(tx)
		goodsInfo := &goods_model.GoodsInfos{
			MerchantId:  merchant.MerchantId,
			Name:        in.GoodsName,
			GoodsTypeId: in.GoodsTypeId,
			Status:      in.Status,
			Img:         in.Img,
			Description: func() string {
				if in.Describe == nil {
					return ""
				}
				return *in.Describe
			}(),
		}
		if err := goodsModel.CreateGoodsInfos(goodsInfo); err != nil {
			return err
		}
		//添加商品的规格
		specificationModel := specification_model.NewSpecificationModel(tx)
		specifications := make([]*specification_model.SpecificationInfos, 0, len(in.SpecificationList))
		for _, specification := range in.SpecificationList {
			specifications = append(specifications, &specification_model.SpecificationInfos{
				GoodsId:          goodsInfo.GoodsId,
				Size:             specification.Size,
				DisplayInventory: specification.DisplayInventory,
				RealInventory:    specification.RealInventory,
				Score:            specification.Score,
			})
		}
		if err := specificationModel.CreateSpecifications(specifications...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("添加失败")
	}

	return &shop.GoodsAddResp{
		OK: true,
	}, nil
}
