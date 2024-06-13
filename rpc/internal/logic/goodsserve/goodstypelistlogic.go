package goodsservelogic

import (
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
	// todo: add your logic here and delete this line

	return &shop.GoodsTypeListResp{}, nil
}
