package logic

import (
	"context"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsAddLogic {
	return &GoodsAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsAddLogic) GoodsAdd(req *types.GoodsAddReq) (resp *types.GoodsAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
