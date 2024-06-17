package logic

import (
	"context"

	"TongChi_shop/api/goods_api/internal/svc"
	"TongChi_shop/api/goods_api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
