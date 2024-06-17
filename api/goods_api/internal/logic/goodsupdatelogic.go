package logic

import (
	"context"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsUpdateLogic {
	return &GoodsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsUpdateLogic) GoodsUpdate(req *types.GoodsUpDateReq) (resp *types.GoodsUpDateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
