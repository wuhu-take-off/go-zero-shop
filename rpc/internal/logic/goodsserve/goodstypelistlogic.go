package goodsservelogic

import (
	"TongChi_shop/model/goods_type_model"
	"context"

	"TongChi_shop/rpc/internal/svc"
	"TongChi_shop/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodsTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTypeListLogic {
	return &GoodsTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品类型列表
func (l *GoodsTypeListLogic) GoodsTypeList(in *shop.GoodsTypeListReq) (*shop.GoodsTypeListResp, error) {
	goodsTypes := goods_type_model.NewGoodsTypeModel(l.svcCtx.DB).FindGoodsTypes()
	var list []*shop.GoodsTypeList
	for i := 0; i < len(goodsTypes); i++ {
		list = append(list, goodsTypes[i].ToGoodsTypesResp())
	}
	return &shop.GoodsTypeListResp{
		GoodsTypeList: list,
		Count:         int32(len(goodsTypes)),
	}, nil
}
