package logic

import (
	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"
	"TongChi_shop/rpc/shop"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsDelLogic {
	return &GoodsDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsDelLogic) GoodsDel(req *types.GoodsDelReq) (resp *types.GoodsDelResp, err error) {

	res, err := l.svcCtx.GoodsRpc.GoodsDel(l.ctx, &shop.GoodsDelReq{
		GoodsId: req.GoodsId,
	})
	if err != nil || !res.OK {
		return nil, err
	}
	return
}
