package logic

import (
	"context"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsListLogic {
	return &GoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsListLogic) GoodsList(req *types.GoodsListReq) (resp *types.GoodsListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
