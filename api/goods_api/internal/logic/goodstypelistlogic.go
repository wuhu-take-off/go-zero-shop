package logic

import (
	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"
	"TongChi_shop/rpc/shop"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTypeListLogic {
	return &GoodsTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTypeListLogic) GoodsTypeList(req *types.GoodsTypeListReq) (resp *types.GoodsTypeListResp, err error) {
	typeList, err := l.svcCtx.GoodsRpc.GoodsTypeList(l.ctx, &shop.GoodsTypeListReq{})
	var list []*types.GoodsTypeList
	for _, goodsTypeList := range typeList.GoodsTypeList {
		list = append(list, &types.GoodsTypeList{
			GoodsTypeId:   goodsTypeList.GoodsTypeId,
			GoodsTypeName: goodsTypeList.GoodsTypeName,
		})
	}
	return &types.GoodsTypeListResp{
		Count: resp.Count,
		List:  list,
	}, nil
}
